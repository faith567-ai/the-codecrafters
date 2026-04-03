# PipeTrail (File String Transformer)

This is a Go CLI project that reads text from a file, applies different string transformation rules, and writes the result to another file.

# Features

The program reads from input.txt, processes each line, and writes the result to output.txt.

Transformations Applied
Trims spaces from each line
Replaces "TODO" with "Action"
Converts:
uppercase lines → title case
lowercase lines → uppercase
Removes empty or blank lines

# How It Works

The program reads all lines from input.txt.
Each line is passed through the pipetrail function.
The function applies different rules depending on the content:
If the line contains "TODO", it replaces it with "Action"
If the line is uppercase, it converts it to title case
If the line is lowercase, it converts it to uppercase
After processing:
Empty lines are removed
The final result is written to output.txt.

# Program Output

When you run the program, it shows:

Number of lines read
Number of lines written
Number of lines removed
Rules applied