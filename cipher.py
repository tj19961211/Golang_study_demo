#!/usr/bin/python

import time
import os.path
from pysqlcipher3 import dbapi2 as sqlcipher


class Timer:
    """Helper class for measuring elapsed time"""
    def __init__(self):
        self.start = None

    def measure(self):
        finish = time.perf_counter()
        elapsed = None if self.start is None else finish - self.start
        self.start = finish
        return elapsed


def make_databse(db_name, key):
    """Create a small database with two tables: version and account.
    Populate tables with some data.
    """
    if os.path.exists(db_name):
        print("Databse {} already exists".format(db_name))
        return
    db = sqlcipher.connect(db_name)
    with db:
        if key:
            db.executescript('pragma key="{}";'.format(key))

        db.execute("CREATE TABLE version(key INTEGER PRIMARY KEY ASC, ver);")
        db.execute("INSERT INTO version(ver) VALUES ('aaa');")

        db.execute("CREATE TABLE account(key INTEGER PRIMARY KEY ASC, name);")
        cur = db.cursor()
        for n_id in range(100):
            cur.execute("INSERT INTO account(name) VALUES ('name {}');".format(n_id))

    print("Test database created: {}, {}".format(
        db_name,
        "Key len={}".format(len(key)) if key else "Not encripted"))

def test_connect_and_reads(run_id, db_name, key, *tables_names):
    """Main test method: connect to db, make selects from specified
    tables, measure timings
    """
    print("{}: Start! Db: {}, {}".format(
        run_id, db_name,
        "Encripted, key len={}".format(len(key)) if key else "Not encripted"))
    timer = Timer()
    timer.measure()

    db = sqlcipher.connect(db_name)
    print("{}: Connect. Elapsed: {} sec".format(run_id, timer.measure()))
    if key:
        db.executescript('pragma key="{}";'.format(key))
        print("{}: Provide Key. Elapsed: {} sec".format(run_id, timer.measure()))
    else:
        print("{}: Skip Provide Key. Elapsed: {} sec".format(run_id, timer.measure()))

    for table_name in tables_names:
        curs = db.execute("SELECT * FROM {};".format(table_name))
        recs = [x for x in curs]
        print("{}: Read {} records from table '{}'. Elapsed: {} sec".format(
            run_id, len(recs), table_name, timer.measure()))

    print("{}: done.".format(run_id))
    print()


def main():
    key = "qwer21"
    #make_databse("rabbits_enc.sqlite3", key)  # prepare encrypted database
    #make_databse("rabbits.sqlite3", "")  # prepare plaintext database

    # test encrypted db
    test_connect_and_reads(0, "log.db", key, 'version', 'account')
   # test_connect_and_reads(1, "rabbits_enc.sqlite3", key, 'account', 'version')
   # test_connect_and_reads(2, "rabbits_enc.sqlite3", key, 'account', 'account')

    # test plaintext db
   # test_connect_and_reads(3, "rabbits.sqlite3", "", 'version', 'account')
   # test_connect_and_reads(4, "rabbits.sqlite3", "", 'account', 'version')
   # test_connect_and_reads(5, "rabbits.sqlite3", "", 'account', 'account')


if __name__ == '__main__':
    main()
