import csv
import random
import uuid
from faker import Faker

# Initialize Faker for generating mock data
fake = Faker()

# Output CSV file name
output_file = "assets/mock_data.csv"

# Column names
columns = [
    "User Id", "Partner Id", "Bond Id", "Ledger Id", "Wallet Account", 
    "Currency", "Transaction Sub Type", "Credit Amt", "Debit Amt", 
    "Remarks", "Requested By", "Assigned Admin name", "Req Date & Time"
]

# Number of records
num_records = 3_000_000

# Sample data for specific columns
currencies = ["USD", "EUR", "GBP", "JPY", "INR"]
transaction_sub_types = ["Purchase", "Transfer", "Refund", "Deposit", "Withdrawal"]
remarks = ["Approved", "Pending", "Rejected", "Completed"]

# Function to generate a single record
def generate_record():
    return {
        "User Id": str(uuid.uuid4()),
        "Partner Id": str(uuid.uuid4()),
        "Bond Id": fake.random_int(min=1000, max=9999),
        "Ledger Id": fake.random_int(min=10000, max=99999),
        "Wallet Account": fake.bban(),
        "Currency": random.choice(currencies),
        "Transaction Sub Type": random.choice(transaction_sub_types),
        "Credit Amt": round(random.uniform(100.0, 10000.0), 2),
        "Debit Amt": round(random.uniform(100.0, 10000.0), 2),
        "Remarks": random.choice(remarks),
        "Requested By": fake.name(),
        "Assigned Admin name": fake.name(),
        "Req Date & Time": fake.date_time_this_year().strftime("%Y-%m-%d %H:%M:%S")
    }

# Generate and write records to CSV
def generate_csv():
    with open(output_file, mode="w", newline="") as file:
        writer = csv.DictWriter(file, fieldnames=columns)
        writer.writeheader()  # Write column headers

        for i in range(num_records):
            writer.writerow(generate_record())
            if (i + 1) % 100_000 == 0:
                print(f"{i + 1} records written...")

    print(f"CSV generation complete. File saved as {output_file}")

# Run the script
if __name__ == "__main__":
    generate_csv()
