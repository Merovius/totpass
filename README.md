# pass-wrapper to generate TOTP

This is a simple wrapper around [pass][pass], the
GPG-based password store, to generate time-based one-time passwords. It expects
Base32-encoded TOTP secrets in the `totp/` folder of your password store. These
secrets are what are encoded in the QR-code you can scan with Google
Authenticator or compatible Apps. If your authentication provider gives you a
way to manually insert the secret, that will also be such a secret. Example usage:

```
$ go get -u github.com/Merovius/totpass
$ pass insert totp/foobar
Enter password for totp/foobar: <paste secret token>
Retype password for totp/foobar: <paste secret token again>
[master 154f06f] Add given password for totp/foobar to store.
$ totpass foobar
431559
$
```

# FAQ

## Why?

I am using Google Authenticator and was annoyed by repeatedly having to
re-register all my 2FA secrets after getting a new phone. I also actually want
the convenience of being able to copy-paste tokens and access my accounts
without my phone available.

## Doesn't this defeat the purpose of 2FA?

Yes and No. In my specific use case, I use a [YubiKey][yubikey] as a GPG key to
encrypt the password store, so there is still essentially a second factor. If
someone exploits my computer, they can still grab the raw secrets during usage.
But if someone exploits my phone, that's the same situation. And in general, I
consider my phone as less trusted than my computers (in particular, if my
computer gets owned, you can already access a lot more interesting information
than my Twitter account).

So, while there might be a *slight* reduction in security, the attack scenarios
I care about (basically phishing and password resets via social engineering)
don't suffer significantly.

## Can this support $thing?

Probably not. This is scratching an itch I am having and I don't think it's
worth maintaining as much of a thing. The good news is, that it's roughly 50
lines of code, so it should be trivial to understand and change it.

# License

```
Copyright 2019 Axel Wagner

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[pass]: https://www.passwordstore.org/
[yubikey]: https://www.yubico.com/products/yubikey-hardware/
