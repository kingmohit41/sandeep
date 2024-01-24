import boto3
import logging
import pandas as pd
from io import BytesIO
import pymongo
from pymongo import MongoClient


def read_s3_data(aws_access_key_id,aws_secret_access_key,bucket_name,prefix,conn_str,database,collection):
    client = MongoClient(conn_str)
    db = client[database] 
    collection = db[collection]
    s3_client = boto3.client('s3', aws_access_key_id=aws_access_key_id,
                            aws_secret_access_key=aws_secret_access_key)


    objects = s3_client.list_objects(Bucket=bucket_name, Prefix=prefix)


    for obj in objects.get('Contents', []):
        file_key = obj['Key']


        if not file_key.endswith('/'):

            response = s3_client.get_object(Bucket=bucket_name, Key=file_key)
            file_content = response['Body'].read()


            df = pd.read_csv(BytesIO(file_content))
            records = df.to_dict(orient='records')
            
            collection.insert_many(records)
            print(f"Inserted {df.shape[0]} rows  into MongoDB")
            logging.info(f"Inserted {df.shape[0]} rows  into MongoDB")

            print(f"File: {file_key}, Shape: {df.shape}")
            logging.info(f"File: {file_key}, Shape: {df.shape}")
            
def funs():
    return "Yes"
