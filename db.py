import psycopg2

def make_conn(args):
    """
    Helper to build up the psycopg2 connection string
    """
    info = []

    if args.dbname:
        info.append("dbname=%s" % args.dbname)
    if args.user:
        info.append("user=%s" % args.user)
    if args.password:
        info.append("password=%s" % args.password)
    if args.port:
        info.append("port=%s" % args.port)

    return ' '.join(info)


def nothing_to_return(sql):
    """
    Evaluate if sql statement will return data or not
    """
    return True if 'CREATE' in sql or 'UPDATE' in sql or 'INSERT' in sql or 'DELETE' in sql else False


def execute(info, sql, args=None, return_columns=False):
    """
    General postgresql execute routine
    """
    with psycopg2.connect(info) as con:
        with con.cursor() as c:
            c.execute(sql, args)

            # don't attempt to return anything if not a SELECT statement
            msg = c.statusmessage
            if nothing_to_return(msg):
                return

            rows = c.fetchall()

            # return the column names with the data
            return tuple([desc[0] for desc in c.description]) + rows if return_columns else rows
