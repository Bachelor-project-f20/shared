# Shared
This repo contains all the model decriptions in the system. It is used as a git submodule in all the other repo's that need to use these models. 

It's added to a repo as a submodule:
```bash
$ git submodule add https://github.com/Bachelor-project-f20/shared
```

And it is updated in the repo's that uses it:
```bash
$ cd shared
$ git fetch
$ git merge origin/master
```
