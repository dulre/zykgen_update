# Zyxel VMG8823, VMG8825 WPA Keygen

https://medium.com/@lucio.corsa/reversing-zyxel-vmg8823-b50b-wpa-algorithm-generation-for-fun-e926e45bf8f3

## Affected SSID
```
"Home&Life SuperWiFi-xxxx" (Password length 16 charaters)
"Infostrada-xxxxxx" (Password length 10 charaters)
"TISCALI_xxxx" (Password length 10 charaters)
```


## Problem

This routers generate their default WPA password using their S/N (serial number), but this number can be easly guessed becouse have got an easy pattern.

### Example

```
"Home&Life SuperWiFi-xxxx"/"Home&Life 2.4GHz" Range:
S182V00000000-S182V99999999
S192V00000000-S192V99999999

"Infostrada-xxxxxx" Range:
S172V00000000-S172V99999999
S182V00000000-S182V99999999

"TISCALI_****" Range:
S172V00000000-S172V99999999
S182V00000000-S182V99999999
```

### Serial Number Algorithm
- The first charaters is always an 'S'.
- The second and third characters should be the year of production ( 2017/2018/2019 for the examples above).
- The string "2V" should be costant for the SSID above, it may be the family's model identifier.
- The rest of the string is different from router to router, but is possibile to generate a dictionary for all the combination to use in combination  with Hashcat.

### List of Zyxel router's
 [Here] (https://pastebin.com/guGVzeNj)

## Usage
'''
zykgen (-m|-n|-c) [-l <length> -L <letter>] <startserial> <endserial> 
'''
Using the command above will create a dictionary in the root directory with all the combination of WPA passwords of the inpunt range, using the specified algorithm, the default letter is the 'V' (which is the fourth letter), while the first 'S' is added automatically and you don't need to insert it, just skip it.
  
### Examples
'''
zykgen.exe -c -l 10 182000000000 182000000010
'''
Will create a dictionary of the WPA passwords with 10 charaters long of the serials ranging from S182V00000000 to S182V00000010 using the 'cosmopolitan' algorithm 

### Type of WPA algorithms 
-*Cosmopolitan* generate WPA password with this combination of charaters: [0-9][A-Z]
-*Negroni* generate WPA password with this combination of charaters: [0-9][A-Z]
-*Mojito* generate WPA password with this combination of charaters: [0-9][A-Z][a-z]
For example the routers above use the Cosmopolitan algorithm, but other router could use another one. 

In order to find which algotithm you router use just look at the back of your router, or if you are unable to do it use google to find images of the back.

## Disclaimer

Don't use it for illegal purpose, this project is mainly used for research purpose and I'm not responsable of it.
