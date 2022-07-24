# go-analyzers
Static analyzers for Go to enforce code consistency and discourage bad practices.

## visibilityorder analyzer

Breaking code example:

![image](https://user-images.githubusercontent.com/1499307/171998356-2f976920-cab5-48b7-8b9a-ec03eb7d5ee5.png)

Resulting compilation error:

![image](https://user-images.githubusercontent.com/1499307/171998390-9e413a54-c84b-4379-b00b-854f5be72a7b.png)


## onlyany analyzer

Forces the use of `any` (introduced in Go 1.8) instead of `interface{}`.
