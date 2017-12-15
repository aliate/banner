# Banner
make banner

## usage

```
fun main() {
    NewBanner("Spring", false).Show()
}
```

it gives:

```
  ___             __ 
 / __| ___   ___ /_/ __    ____
\__ \ / _ \ / __\/ // _ \ / _ / 
|___// ___//_/  /_//_//_/_\_ / 
    /_/                  \__/
```

> `adjoin` mode:

```
fun main() {
    NewBanner("Spring", true).Show()
}
```

it shows:

```
  ___           __          
 / __|___  ___ /_/__    ____
\__ \/ _ \/ __\/ / _ \ / _ /
|___/ ___/_/  /_/_//_/_\_ / 
   /_/                \__/ 
```