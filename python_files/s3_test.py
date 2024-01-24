import boto3
import pandas as pd
import logging


def read_s3():
    aws_access_key_id = 'AKIARV7DJ7DWPA75Y2LZ'
    aws_secret_access_key = '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'
    bucket_name = 'athenabdb'
    file_key = f"job/conversation_attributes_202305101307.csv"

    s3 = boto3.client('s3', aws_access_key_id=aws_access_key_id, aws_secret_access_key=aws_secret_access_key)

    obj = s3.get_object(Bucket=bucket_name, Key=file_key)
    df = pd.read_csv(obj['Body'])
    logging.info(df)
    return df

