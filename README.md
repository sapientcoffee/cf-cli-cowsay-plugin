# CF CLI Cowsay
This is a fun learning exercise to teach me how to create a cf cli and also help teach me go.

I am sure that this could be written in a more efficient way however for the purpose of learning it is what it is. My goal was to have a working `cf cli` plugin by writing my 1st golang progam that is testable (new concept as well).

### Still to do

* Extract the app details and status for space and display
* Complete the tests with ginkgo
* Complete this Readme page

## Usage

```
➔ cf cowsay
 ________
< oh hey >
 --------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

```
➔ cf cowsay hello
 _______
< hello >
 -------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

```
➔ cf cowsay space
 _________________________________________
< Space: rob in the organisation: testing >
 -----------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

```
➔ cf cowsay apps

WIP: but will display the apps in the space and thier status.
```

### Install

`cf install-plugin cowsay`


### Delete 
`cf uninstall-plugin cowsay`

## Working with the code

### Testing

### Compile


