# Zyxel VMG8823, VMG8825 WPA Keygen

https://medium.com/@lucio.corsa/reversing-zyxel-vmg8823-b50b-wpa-algorithm-generation-for-fun-e926e45bf8f3

## Affected SSID
'''
"Home&Life SuperWiFi-xxxx" (Password length 16 charaters)
"Infostrada-xxxxxx" (Password length 10 charaters)
"TISCALI_xxxx" (Password length 10 charaters)
'''


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

### Usage
'''
zykgen (-m|-n|-c) [-l <length> -L <letter>] <startserial> <endserial>
'''

## Disclaimer

This project is licensed under the MIT License.
