import random,logging
import pandas as pd
def emp_data():
    l=[]
    for i in range(2000):
        d={"emp_id":i, "department":random.choice(['IT','HR','Sales','Operations']),
          "salary":random.randint(10000,100000),"working_mode":random.choice(['Hybrid','onsite'])}
        logging.info(d)
        l.append(d)
    df=pd.DataFrame(l)
    return df
