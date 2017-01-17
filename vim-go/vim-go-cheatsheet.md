### Cheat sheet for vim-go

#### Build Run Test Import
- go build
```
:GoBuild
,b
```

- go run
```
:GoRun
,r
```

- go test
```
:GoTest
,t
```

- go test func
```
:GoTestFunc
```

- go coverage
```
:GoCoverage
:GoCoverageClear
:GoCoverageToggle
,c
```

- go import
```
:GoImport strings
:GoDrop strings
```

#### Inside Function
- delete inside of a function, the `if` stands for inside of a function
```
dif
```

- select text inside of a function
```
vif
```

- yank text inside of a function
```
yif
```
#### All Function
- Use `af` any where in the function to select, delete and yank all the function, including comments, you can disable comments if you would like to in the vimrc
```
daf
vaf
yaf
```
