from azure.storage.blob import BlobServiceClient
import os

def upload_file_to_blob(storage_account_url, container_name, blob_name, file_path, connection_string):
    blob_service_client = BlobServiceClient.from_connection_string(connection_string)
    blob_client = blob_service_client.get_blob_client(container=container_name, blob=blob_name)
    
    with open(file_path, "rb") as data:
        blob_client.upload_blob(data, overwrite=True)
    
    print(f"Uploaded {file_path} to blob storage: {blob_name}")

def download_blob_to_memory(storage_account_url, container_name, blob_name, connection_string):
    blob_service_client = BlobServiceClient.from_connection_string(connection_string)
    blob_client = blob_service_client.get_blob_client(container=container_name, blob=blob_name)
    
    download_stream = blob_client.download_blob()
    return download_stream.readall()
