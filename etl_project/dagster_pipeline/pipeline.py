from dagster import pipeline, solid, Field
from dagster_pipeline.transformations.product_transform import transform_and_load
from azure_blob.blob_operations import download_blob_to_memory

@solid(config_schema={"storage_account_url": Field(str), "container_name": Field(str), "blob_name": Field(str), "connection_string": Field(str)})
def download_blob(context):
    config = context.solid_config
    blob_data = download_blob_to_memory(config["storage_account_url"], config["container_name"], config["blob_name"], config["connection_string"])
    return blob_data

@solid(config_schema={"storage_account_url": Field(str), "container_name": Field(str), "blob_name": Field(str), "connection_string": Field(str), "elasticsearch_host": Field(str), "elasticsearch_index": Field(str)})
def transform_and_load_data(context, blob_data):
    config = context.solid_config
    transform_and_load(blob_data, config["elasticsearch_host"], config["elasticsearch_index"])

@pipeline
def etl_pipeline():
    blob_data = download_blob()
    transform_and_load_data(blob_data)
