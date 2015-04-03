import unittest
import argparse
import db

def create_args():

    args = argparse.Namespace
    args.dbname = None
    args.user = None
    args.password = None
    args.port = None
    return args


class TestDB(unittest.TestCase):

    def test_make_conn(self):

        args = create_args()
        actual = ""
        info = db.make_conn(args)
        self.assertEqual(actual, info)
    
        args = create_args()
        args.dbname = "dumptruck"
        actual = "dbname=dumptruck"
        info = db.make_conn(args)
        self.assertEqual(actual, info)
    
        args = create_args()
        args.user = "dev"
        actual = "user=dev"
        info = db.make_conn(args)
        self.assertEqual(actual, info)
    
        args = create_args()
        args.password = "test"
        actual = "password=test"
        info = db.make_conn(args)
        self.assertEqual(actual, info)
    
        args = create_args()
        args.port = 5432
        actual = "port=5432"
        info = db.make_conn(args)
        self.assertEqual(actual, info)
    
        args = create_args()
        args.dbname = "dumptruck"
        args.port = 5432
        actual = "dbname=dumptruck port=5432"
        info = db.make_conn(args)
        self.assertEqual(actual, info)
    
    def test_nothing_to_return(self):

        sql = "CREATE 1"
        self.assertEqual(db.nothing_to_return(sql), True)
        sql = "UPDATE 1"
        self.assertEqual(db.nothing_to_return(sql), True)
        sql = "INSERT 1"
        self.assertEqual(db.nothing_to_return(sql), True)
        sql = "DELETE 1"
        self.assertEqual(db.nothing_to_return(sql), True)
        sql = ""
        self.assertEqual(db.nothing_to_return(sql), False)
    
    def test_execute(self):
        pass
    
    def test_executemany(self):
        pass

if __name__ == '__main__':
    unittest.main()
