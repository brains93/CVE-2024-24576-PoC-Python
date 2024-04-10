import subprocess

def main():
    print("enter payload here")
    input_payload = input()
    try:
        output = subprocess.run([r"C:\Users\Hacker\source\repos\CVE-2024-24567-PoC-Python/test.bat", input_payload], capture_output=True, shell=True, text=True)
        print("Output:\n", output.stdout)
        print("Errors:\n", output.stderr)  # Print any errors
    except subprocess.CalledProcessError as e:
        print("Error:", e)

if __name__ == "__main__":
    main()
