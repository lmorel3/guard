# Guard

<p align="center">
  <img src="https://raw.githubusercontent.com/lmorel3/guard/master/assets/screenshot.png">
</p>

**Guard** is an open-source _**simple and lightweight**_ SSO authentication handler for reverse proxies, written in **Go**.

**Guard** aims to make an easily configurable SSO handler, which works with various reverse proxies.
**Guard** will stay **simple**, without 2FA, LDAP support, etc. \
If you want these features, have a look at [Authelia](https://github.com/clems4ever/authelia) ;)

Currently supported reverse proxies:

- [Traefik](https://traefik.io/)
- _Every reverse proxies which forward authentication via X-Forwarded-* headers_
- _Coming soon if you want to contribute :)_

## Getting started
You can have a try using the example.

1. First, edit _/etc/hosts_ and add the following lines:
```
127.0.0.1       guard.local
127.0.0.1       auth.guard.local
127.0.0.1       public.guard.local
```

2. Then, simply go to `example` folder and run `docker-compose up` 

3. Open a browser and navigate to `http://guard.local`: you should be redirected to `https://auth.guard.local`.

4. Use the default credentials `admin`/`admin` so that you are redirected to the app. You're now logged in!

5. If you logout, you should be able to access to `http://public.guard.local` which is publicly allowed.

## Configuration

The configuration is _minimalist_. Simply provide your domain, the subdomain used for **Guard** and possible public URLs.

```
domain: guard.local
guard: auth.guard.local
allowed:
  - public.guard.local
```

_Note:_ URL can be more precise (e.g. `xyz.guard.local/public`)

## Administration

<p align="center">
  <img src="https://raw.githubusercontent.com/lmorel3/guard/master/assets/screenshot_admin.png">
</p>

**Guard** provides an basic and easy-to-use admin interface, in which you can add or remove users.

## Thanks
- [xyproto/permissionbolt](https://github.com/xyproto/permissionbolt)
- [lmorel3/guard-php](https://github.com/lmorel3/guard-php) : Guard was initially written in PHP, but I decided to use **Go** for better performances!