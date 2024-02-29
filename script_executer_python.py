import random,logging
import pymongo
from pymongo import MongoClient
import logging

def emp_data():
    l=[]
    for i in range(20):
        d={"emp_id":i, "department":random.choice(['IT','HR','Sales','Operations']),
          "salary":random.randint(10000,100000),"working_mode":random.choice(['Hybrid','onsite'])}
        logging.info(d)
        l.append(d)
    return l


def write(l,conn_str):
    client = pymongo.MongoClient(conn_str)
    db = client["pipeline"] 
    collection = db["emp_data"]  
    collection.insert_many(l)

    
def start(conn_str):
    l=emp_data()
    logging.info("=============================Employee data read===============================")
    w=write(l,conn_str)
    logging.info("=============================Data Written=====================================") 
    return ["Data Read and written successfully."]