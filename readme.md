# Injector

**Injector** is a simple Go program designed to inject one or more executable files into a single binary executable. The injected executables are embedded as byte slices within the resulting binary, allowing the combined binary to execute each of the injected executables when run.

## How it Works

Injector uses the following steps to accomplish its task:

1. Determine the operating system (Windows, Linux, or macOS) to appropriately handle the binary file extension.
2. Parse command-line arguments to identify the target name for the resulting binary and the names of the executables to inject.
3. Read the content of each executable to inject and store them as byte slices.
4. Create a wrapping Go file containing embedded byte slices for each executable.
5. Build a new binary by compiling the wrapping Go file and the target name.
6. Clean up temporary wrapping Go file.

## Usage

To use Injector, follow these steps:

1. Make sure you have Go installed on your system.
2. Download the "injector.go" file and place it in a directory of your choice.
3. Open a terminal or command prompt and navigate to the directory containing "injector.go."
4. Either compile it or just run the code

To compile:

    Compile the "injector.go" file with the following command:

    ```
    go build injector.go
    ```

    Run the compiled "injector" binary with the following format:

    ```
    ./injector <target_name> <executable_name_1> <executable_name_2> ... <executable_name_n>
    ```

    - `<target_name>`: The name of the resulting binary after injection.
    - `<executable_name_1> <executable_name_2> ... <executable_name_n>`: The names of the executables you want to inject into the target binary.

To run:

    ```
    go run injector.go <target_name> <executable_name_1> <executable_name_2> ... <executable_name_n>
    ```

5. If a file with the `<target_name>` already exists in the current directory, Injector will prompt you to confirm whether you want to overwrite it.

6. Once the process is complete, you will have a new binary named `<target_name>` in the current directory. Running this binary will execute each of the injected executables in sequence.
    ```
    ./<target_name>
    ```
7. Two simple hello world programs are provided, so an example use case would be:
    ```
    gcc files/helloc.c -o helloc
    go build files/hellogo.go
    go run injector.go hello helloc hellogo
    ```

## Note

1. **Windows Executable Extension**: For Windows, the injector will automatically append the ".exe" extension to the `<target_name>` to create an executable with the appropriate extension.

2. **Go Dependencies**: The injector utilizes the "github.com/google/uuid" and "github.com/amenzhinsky/go-memexec" packages for generating a temporary name and executing the injected executables. These packages will be automatically downloaded during the compilation process.

3. **Supported OS**: The injector currently supports Windows, Linux, and macOS. Running the injector on an unsupported OS will result in an error.

4. **File Existence**: Ensure that all executable files to be injected exist in the current directory or have the correct relative/absolute paths.

## Disclaimer

The injector program is provided as-is without any warranty. Please use it responsibly and avoid using it for malicious purposes or without the consent of the involved parties.
