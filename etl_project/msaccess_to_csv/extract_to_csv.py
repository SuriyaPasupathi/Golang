import pandas as pd
import pyodbc

def extract_data_from_msaccess(database_path: str, query: str, output_csv_path: str):
    conn_str = f'DRIVER={{Microsoft Access Driver (*.mdb, *.accdb)}};DBQ={database_path};'
    conn = pyodbc.connect(conn_str)
    
    df = pd.read_sql(query, conn)
    df.to_csv(output_csv_path, index=False)
    
    conn.close()
    print(f"Data extracted to {output_csv_path}")
