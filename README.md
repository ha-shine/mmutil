# Myanmar Unicode Utilities

```mmutil``` is a small golang utility package for myanmar unicode and zawgyi.

# Supported Functions
## func **StringIsZawgyi**

```func StringIsZawgyi(s string) bool```

StringIsZawgyi detects whether the given string is written in Zawgyi or not.

## func **IsZawgyi**

```func IsZawgyi(b byte[]) bool```

IsZawgyi detects whether the given byte array is written in Zawgyi or not.

## func **SplitWords**

```func SplitWords(s string) []string```

SplitWords split given unicode string into individual words. e.g

```go
result := SplitWords("ကိုကြီး") // ["ကို", "ကြီး"]
```

# Todos
- Conversion between zawgyi and unicode (Currently available converters use Regular expressions with back reference. Golang doesn't support back reference.)
- Support kinzi, and english words in SplitWords function. (kinzi - "င်္")