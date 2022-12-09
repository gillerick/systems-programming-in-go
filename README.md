## Systems Design

This project is based on the study of Mihalis
Tsoukalos' [Go Systems Programming](https://www.packtpub.com/product/go-systems-programming/9781787125643)

### Regular Expressions

| Metacharacter     | Meaning                                                                                                                                                                              |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ^                 | Matches the beginning of the line only                                                                                                                                               |
| $                 | Matches the end of the line only                                                                                                                                                     |
| *                 | A single character followed by an * will match zero or more occurrences                                                                                                              |
| []                | Character enclosed inside the [] will be matched. This can be a single or range of characters. You can use the - to include a range inclusively. Instead of saying [12345] use [1-5] |
| \                 | This is used to escape the special meaning of a metacharacter                                                                                                                        |
| .                 | Matches any single character                                                                                                                                                         |
| _pattern \{n \}_  | Matches a specific number of occurrences specified by n that contains the preceding character                                                                                        |
| _pattern \{n,\}m_ | Works like the above pattern but matches at least n occurrence if the preceding pattern                                                                                              |
| _pattern \{n,m\}_ | Matches any number of occurrences between n and m of the preceding pattern                                                                                                           |


## wc
