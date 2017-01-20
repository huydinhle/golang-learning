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

#### Split join structs 
- multiple lines spliting
```
gS
```
- Join all the lines together
```
gJ
```

#### snippets
- Incredibly powerful tools and There are a shit tons to learn
```
errp<tab>
ff<tab>
lf<tab>
json<tab>
```

### Lint your code
```
:GoMetaLinter
saving by  :w
```

### Go in definition
```
:GoDef
ctrl-]
gd
```

### Go back definition
```
:GoDefPop
ctrl-t
gd
```

### GoDef Stack manipulation
```
:GoDefStack
:GoDefStackClear
```

### Move between functions
```
:GoDecls
:GoDeclsDir
```

### Move to the next function within go files
```
[[
]]
d[[
4[[
```

### Alternate file movement
```
:A --> alternate opening
:AS --> Alternate split horizontally
:AV --> Alternate split vertically
```

### Go documentations
```
:GoDoc
```

### Go Info
```
:GoInfo
,i
```

### Go files and dependencies
```
:GoFiles
:GoDeps
```

### Guru
```
GoReferrers
GoDescribe
GoImplements
GoWhicherrs
GoChannelPeers
```

#### Function tracing in go
```
GoCallees
GoCallers
GoCallstack
```
### Refactor Rename
```
:GoRename bar
```
### Refactor a function out of a piece of cod
- `:GoFreevars` telling you how many variables dependent on the piece of code you looking at
- select the bunc of code you want to put in function, then do `:GoFreevars`

### Generate code
- Implement an interface
```
:GoImpl
io.ReadWriteCloser
```
### Share your code go playground
:GoPlay
