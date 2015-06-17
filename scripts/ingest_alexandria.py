import time
from db import make_conn, execute, executemany
import argparse
import filters

def create_table(name):

    # only create if doesn't exist yet
    if name in filters.getFilterTables(cs):
        return

    print 'Creating "%s" filter table..' % name

    sql = '''
    INSERT INTO filter_descriptions(name, display_name, data_type) 
    SELECT '%s', 'Alexandria Sentiment Data', 'MULTI' 
    WHERE NOT EXISTS (
        SELECT name FROM filter_descriptions WHERE name = '%s' LIMIT 1
    );
    ''' % ((name,) * 2)
    execute(sql, info=cs)

    sql = '''
    DROP TABLE IF EXISTS %s;
    CREATE TABLE %s (

        -- base columns for filtering table
        sid TEXT COLLATE "C" REFERENCES security_master (sid)
        , day DATE -- DATE(ts)
        , val NUMERIC -- the default filtering value
        , UNIQUE(ts, sid)

        -- plus extras that can be specified for filtering
        , storyid TEXT
        , ts TIMESTAMP
        , country TEXT
        , sentiment INT
        , confidence NUMERIC
        , novelty INT
        , subjects TEXT
        , relevance NUMERIC
    );
    CREATE INDEX ON %s (sid, val);
    ''' % ((name,) * 3)
    execute(sql, info=cs)


def parse_line(line):

    StoryId, Timestamp, Ticker, Country, Sentiment, Confidence,  \
    Novelty, Subjects, Relevance = line.split('\t')

    default = Sentiment

    args = (Timestamp, default, StoryId, Country, Sentiment, \
    Confidence, Novelty, Subjects, Relevance, Ticker, Timestamp, \
    Ticker, Timestamp, default, StoryId, Timestamp, Country, \
    Sentiment, Confidence, Novelty, Subjects, Relevance)

    return args


def parse(fname):

    name = 'alexandria'
    create_table(name)

    lines = (line for line in open(fname, 'r'))
    lines.next()

    sql = '''
    WITH upsert AS (
        UPDATE ''' + name + ''' 
            SET day = DATE(%s)
            , val = %s
            , storyid = %s
            , country = %s
            , sentiment = %s
            , confidence = %s
            , novelty = %s
            , subjects = %s
            , relevance = %s
        WHERE 
            sid = (SELECT sid FROM security_master WHERE symbol = %s LIMIT 1)
            AND ts = %s
        RETURNING *
    ) 
    INSERT INTO ''' + name + ''' (
        sid, day, val, storyid, ts, country, sentiment, confidence, novelty,
        subjects, relevance )
    SELECT 
        (SELECT sid FROM security_master WHERE symbol = %s LIMIT 1)
        , DATE(%s), %s, %s, %s, %s, %s, %s, %s, %s, %s
    WHERE NOT EXISTS (SELECT 1 FROM upsert);
    '''

    batch_size = 10000
    args = []
    for line in lines:
        arg = parse_line(line)
        args.append(arg)
        if len(args) % batch_size == 0:
            start = time.time()
            executemany(sql, args, info=cs)
            print 'Submitted %d rows in %f seconds' % (batch_size, time.time() - start)
            args = []


if __name__ == '__main__':

    parser = argparse.ArgumentParser()
    parser.add_argument(
        '--host', default='localhost', help='Database host name')
    parser.add_argument(
        '--dbname',  default='invest', help='Database dbname name')
    parser.add_argument(
        '--user',  default='dbadmin', help='Database user name')
    parser.add_argument(
        '--password',  help='Database password')
    parser.add_argument(
        '--port',  default=5432, type=int, help='Database port name')
    parser.add_argument(
        '--fname', required=True, help='Specific csv to process')
    args = parser.parse_args()

    cs = make_conn(args)
    parse(args.fname)
