# vietnam-zipcode
Takes `zipcode` as input and returns its corresponding `province` and `ward` 

These data were obtained from this site: [Thu vien phap luat](https://thuvienphapluat.vn/chinh-sach-phap-luat-moi/vn/ho-tro-phap-luat/chinh-sach-moi/94242/tra-cuu-ma-buu-chinh-cua-cac-xa-phuong-34-tinh-thanh-sau-sap-nhap). 

If you are using my code, pls consider leaving a star ⭐

Check out this service on Rapid Api: [Vietnam Zipcode](https://rapidapi.com/quanghia24/api/vietnam-zipcode)
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
