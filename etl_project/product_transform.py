import pandas as pd
from io import BytesIO
from elasticsearch import Elasticsearch

def transform_data(blob_data):
    df = pd.read_csv(BytesIO(blob_data))
    # Add transformations here, e.g., renaming columns, filtering data
    df['price_with_tax'] = df['price'] * 1.1
    return df

def load_data_to_elasticsearch(df, es_host, es_index):
    es = Elasticsearch([es_host])
    for _, row in df.iterrows():
        doc = row.to_dict()
        es.index(index=es_index, document=doc)

def transform_and_load(blob_data, es_host, es_index):
    df = transform_data(blob_data)
    load_data_to_elasticsearch(df, es_host, es_index)
    print("Data loaded into Elasticsearch")
