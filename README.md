# A credit card validator RESTful api service written in Golang

## It has two apis, whose enpoints are as below

### **/check_issuer** 
>this checks for card issuer name like Visa/Master Card
>
>this api can be used while the user starts writing the card number


 Request for the above api will be like below 

```
{
     "card_number": "46"
}
```

 Response will be as below


```
If issuer is valid ->

{
    "is_valid_issuer": true,
    "issuer": "Visa"
}

if issuer is not valid

{
    "is_valid_issuer": false,
    "issuer": ""
}
```


###  **/validate** 
>this validate card number with [LUHN or modulo 10 algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm)
>
>this api should be used to verify the entire card number for accidental mistakes by the user


### Request for the above api will be as below 

```
{
     "card_number": "6522940767680906"
}
```

### Response will be as below 

```
if Issuer is fine and card is also valid
{
    "card_number": "6522940767680906",
    "is_valid": true,
    "issuer": "Rupay"
}

If issuer is not fine

"We only accept card from Visa, MasterCard or Rupay"

If issuer is fine but card number if invalid

{
    "card_number": "632294076760906",
    "is_valid": false,
    "issuer": "Rupay"
}
```
