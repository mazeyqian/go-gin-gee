# ESLint Files

This Go script is designed to format a set of files using ESLint. ESLint is a tool for identifying and reporting on patterns found in ECMAScript/JavaScript code, with the goal of making code more consistent and avoiding bugs.

Here's a basic guide on how to use this script.

## Command Line Flags

The script accepts several command line flags:

| Flag | Description | Example | Required |
| --- | --- | --- | --- |
| `files` | A comma-separated list of files to be formatted. | `file1.js,file2.js` | Optional |
| `folders` | A comma-separated list of folders to be formatted. | `src/utils,src/components` | Optional |
| `esConf` | The path to the ESLint configuration file. | `.eslintrc.js` | Optional |
| `esCom` | The ESLint command to run. | `--fix` | Optional |
| `root` | The root of the folders. | `src` | Optional |
| `ext` | The file extension to look for when formatting files. | `.js` | Optional |
| `befCom` | Commands to run before the formatting step. | `echo 'Starting format'` | Optional |
| `aftCom` | Commands to run after the formatting step. | `echo 'Format completed'` | Optional |
| `filesRang` | The range of files to format. | `1-10` | Optional |

1. **File and Folder Selection**: The script allows you to specify the files and folders to be formatted using the `files` and `folders` flags. If a root directory is specified, it will prepend the root to each file or folder. It will also find all files with the specified extension within each folder.
2. **Command Execution**: If any commands are specified in the `befCom` or `aftCom` flags, they will be executed before and after the formatting step, respectively.
3. **ESLint Execution**: For each file in the list, the script will execute the ESLint command with the specified configuration file. It will log any errors that occur during this process.
4. **File Range**: If a file range is specified, the script will find all files within this range with the specified extension. It will then filter out any files that were already formatted, and log the remaining files.
5. **Logging**: The script logs various information throughout its execution, including the number of files worked on, any errors that occur, and any files within the specified range that were not formatted.

## Examples of Usage

To run the script, you would use a command like this:

1\. Simplest command:

```bash
go run eslint-files.go -files="file1.js,file2.js"
```

This command will format the files `file1.js` and `file2.js` using the default ESLint configuration and command.

2\. Complex command:

```bash
go run eslint-files.go -files="file1.js,file2.js" -folders="src/utils,src/components" -esConf=".eslintrc.js" -esCom="--fix" -root="src" -ext=".js" -befCom="echo 'Starting format'" -aftCom="echo 'Format completed'" -filesRang="1-10"
```

This command will format the files `file1.js` and `file2.js`, as well as all `.js` files in the `src/utils` and `src/components` folders. It will use the ESLint configuration in `.eslintrc.js` and the ESLint command `--fix`. It will also echo a start and end message, and only format files in the range 1-10.

You can adjust these parameters as needed to suit your specific use case.
