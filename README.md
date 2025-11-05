# vietnam-zipcode
Takes `zipcode` as input, returns `province` and `ward`

If you are using my code, pls consider leaving a star ⭐

```
@GET /90469

:200
{
    "province": "TỈNH AN GIANG",
    "ward": "X. Vĩnh Hậu"
}
```

```
@GET /69

:404
{
    "error": "zip code not found"
}
```