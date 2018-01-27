What if you wanted to find other people who like cryptocurrencies and spend your money with them rather than just using some random service from the yellow pages? Well, [cryptopages](https://cryptopages.club).

## Authentication

This app makes use of JWT authentication. For this to work on your machine, you need to generate keys. To do this, from the root of the project:

    cd jwtkey
    openssl genrsa -out jwt.key 2048 && openssl rsa -in jwt.key -outform PEM -pubout -out jwt.key.pub
