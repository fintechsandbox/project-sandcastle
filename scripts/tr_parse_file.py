'''
Sample script to parse fundamental values from a 
Thomson Reuter file downloaded from their FTP site. 
Each set of fundamentals is written to its own 
database table using the python [postgres] db script. 
'''

import xml.etree.ElementTree as ET
import sys
import datetime
import time
import multiprocessing
import db
import argparse
import os

# make sure exchange is listed properly 
lookup = { 'NASD' : 'NASDAQ', 'NYSE' : 'NYSE', 'AMEX' : 'AMEX' } 

# output codes
EXT='.parsed'
CURRENCY='.currency' + EXT
EXCHANGE='.exchange' + EXT
FIELD='.field' + EXT
source = 'Thomson Reuters'

# parse a single interim standardized fundamental file
def parse(fname):

    print fname

    # remove any previously appended extensions
    originalName = fname
    while fname.count('.') > 1:
        fname,extension = os.path.splitext(fname)
    os.rename(originalName, fname)

    tree = ET.parse(fname)
    root = tree.getroot()

    # for reference

    repno = root.find('RepNo').text
    description = root.find('ReferenceInformation/CompanyInformation/Company').get('Name')
    exchange = root.find('ReferenceInformation/Issues/Issue/Exchange').get('Code')

    sedol = None
    symbol = None
    ric = None
    for issue in root.findall('ReferenceInformation/Issues/Issue/IssueXref'):
        if issue.get('Type') == 'SEDOL':
            sedol = issue.text
        elif issue.get('Type') == 'Ticker':
            symbol = issue.text
        elif issue.get('Type') == 'DisplayRIC':
            ric = issue.text

    for i in [ repno, description, exchange, sedol, symbol, ric ]:
        if not i:
            print "Missing field.. skipping %s" % fname
            os.rename(fname, fname + FIELD)
            return

    name = None
    if exchange and symbol:
        exch = lookup.get(exchange)
        if not exch:
            print "Unsupported exchange '%s'.. skipping %s" % (exchange, fname)
            os.rename(fname, fname + EXCHANGE)
            return
        else:
            exchange = exch
        name = '%s:%s' % (exchange, symbol)

    currency = root.find('ReferenceInformation/CompanyInformation/ReportingCurrency').get('Code')

    # only updates for now to ensure companies that don't have eoddata aren't inserted
    # sids,exchange,symbol get inserted by ingest_eoddata.py
    buf = '''
	UPDATE security_master 
	SET repno = %s
	, ric = %s
	, sedol = %s
	, display_name = %s
	WHERE exchange = %s AND symbol = %s
    '''
    args = (repno,ric,sedol,description,exchange,symbol)
    db.execute(buf, args)
    
    # only want companies reporting in USD 
    if currency != 'USD':
        print 'Non USD currency field in', fname,'.. skipping'
        os.rename(fname, fname + CURRENCY)
        return

    # delimiter for the statement type + coa combo
    delim = '_'

    # upsert filter_descriptions on filter name 
    buf = '''
    WITH upsert AS (
        UPDATE filter_descriptions
        SET description = %s, source = %s
        WHERE filter = %s
        RETURNING *
    )   INSERT INTO filter_descriptions (filter,description,source)
        SELECT %s,%s,%s
        WHERE NOT EXISTS (SELECT 1 FROM upsert);
    '''
    args = []
    for layout in root.iter('Layout'):
        layoutType = layout.get('Type')
        for mapitem in layout.iter('MapItem'):
            coa = layoutType + delim + mapitem.get('COA')
            coa = coa.upper()
            description = mapitem.text
            args.append((description, source, coa, coa, description, source))

            # create the filter table if doesn't exist
            db.createFilterTable(coa)

    db.executemany(buf, sorted(args)) 

    # periodEnd is most recent date the current statement applies to. 
    # Because we iterate from newest to oldest, we can save the periodStart 
    # from the previous statement and use it as the stop for the current statement.
    # We use today's date as the stop for the most recent statement.
    periodEnd = datetime.date.today().strftime('%Y-%m-%d')

    for i,period in enumerate(root.iter('Period')):

        # only process the most recent period
        if onlyMostRecent and i > 0:
            break

        periodStart = period.get('PeriodEndDate')

        # write one day,val of filter data for each day we have intraday data for
        buf = "SELECT DISTINCT day FROM days WHERE day BETWEEN %s AND %s"
        args = (periodStart, periodEnd)
        rows = db.execute(buf, args, returndata=True)

        # skip periods that don't have corresponding intraday data 
        if not rows:
            continue

        print fname,periodStart,periodEnd
        days = [ x[0] for x in sorted(rows) ]

        for statement in period.iter('Statement'):
            statementType = statement.get('Type')
            for fv in statement.iter('FV'):
                coa = statementType + '_' + fv.get('COA')
                val = fv.text

                # upsert on name,day
                # update the val for this table's name,day or insert if no val for this name,day
                buf = '''
                WITH upsert AS (
                    UPDATE ''' + coa + ''' 
                    SET val = %s 
                    WHERE name = %s AND day = %s
                    RETURNING *
                )
                    INSERT INTO ''' + coa + ''' (name,day,val) 
                    SELECT %s,%s,%s 
                    WHERE NOT EXISTS (SELECT 1 FROM upsert);
                '''
                    
                args = []
                for day in days:
                    args.append((val,name,day,name,day,val))
                # bc processed chronologically, 'Reclassified Normal' entries will be submitted
                # first and should NOT be replaced by 'Update Normal' entries
                db.executemany(buf, args)

        periodEnd = periodStart

    # rename file to help us track what files still need to be processed
    os.rename(fname, fname + EXT)
    print 'Finished',fname

if __name__ == '__main__':

    parser = argparse.ArgumentParser(description='Process data files from TKRD FTP')
    parser.add_argument('--recent', action='store_true', default=False, help='add flag to only process the most recent periods in all statements')
    parser.add_argument('--force', action='store_true', default=False, help="add flag to force reprocessing of files (processed files end in '.parsed')")
    parser.add_argument('--debug', action='store_true', default=False, help='add flag to run in serial, with debug statements')
    parser.add_argument('--files', required=True, nargs='+', help='Files names to process')
    args = parser.parse_args()

    # set to True to only process the most recent periods in all statements
    onlyMostRecent=args.recent

    files = []
    for fname in args.files:

        # don't reprocess files if flag not set
        if fname.endswith(EXT) and args.force == False:
            continue

        files.append(fname)

    if args.debug:
        # parse in serial
        for f in files:
            parse(f)
    else:
        # use python multiprocessing
        # to parse one file per process
        pool = multiprocessing.Pool()
        pool.map_async(parse, files).get(999999)
