

### Option 1 (using source and self_package)
Command 
```
mockgen -package=testdata -destination=./testdata/mocks1.go -source=bar/interfaces.go -self_package=mydemo/bar 
```

### Option 2 (using source and no self_package)
Command 
```
mockgen -package=testdata -destination=./testdata/mocks2.go -source=bar/interfaces.go  
```

### Option 3 (using reflect mode and self_package )
Command 
```
mockgen -package=testdata -destination=./testdata/mocks3.go -self_package=mydemo/bar mydemo/bar DoShomething,MoreStuff
```

### Option 4 (using reflect mode and no self_package)
Command 
```
mockgen -package=testdata -destination=./testdata/mocks4.go mydemo/bar DoShomething,MoreStuff
```