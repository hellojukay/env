# env

```
vagrant@archlinux env (main) $ ./env -h >> README.md
Usage of ./env:
  -i    clear the environment then run command:
  -u value
        remove variable from the environment and run a program:
```

# Use case

1. Show the environment:
```
env
```
2. Clear the environment and run a program:
```
env -i program
```
3.  Remove variable from the environment and run a program:
```
env -u variable program
```
4. Set multiple variables and run a program:
```
env variable1=value variable2=value variable3=value program
```