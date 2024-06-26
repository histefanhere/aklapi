package aklapi

const rootHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Healthcheck</title>
    <link href="https://fonts.googleapis.com/css?family=Oswald&display=swap" rel="stylesheet">
    <style>
    body {
        font-family: 'Oswald', 'Arial Narrow', sans-serif;
        text-align: center;
    }
    </style>
</head>
<body>
    <img src="data:image/png;base64,{{.CyberdyneLogo}}">
    <p>AT SOME POINT IN TIME, FUTURE WILL BE DIFFERENT FROM THE PRESENT.</p>
</body>
</html>
`

const cyberdynePng = `
iVBORw0KGgoAAAANSUhEUgAAAQAAAAC5CAMAAADwHhFdAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAA
B6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAACslBMVEWgoL2ZmcZYWFhxcXF7e3uR
kZTIyMi9vc7S0tI1ODnW1tYqLC1CQkLY2NgMDAxOTk5lZWXNzc01NTUSEhPc3O7U1NSrq9bBwd6BgY
HR0tFJR0aens+qqtVHRkWmpsnIyOTZ2d69vebBwdjGxuhcXFzc3NzKzs9UVFSxsdqoqNZMUla1tdys
rNiiotKnp9W0tNuiotKkpNOXl8udnc+8vOGgoNCoqNWqqta3t967u+CoqNW5ud6trdi3t9ufn9Cvr9
irq9esrNerq9imptScnM6kpNOkpNK3t92vr9rh4uGCgH7U1NRfVEkpJiJRTkwJBAQAAACrq6tqaGaz
s7OUlJSYo6tiYF67u7yZmZlGSUslIR3Av7+kpKSQkI/b29uEhISWl5l0c3JTUU4bGxs9OzotKyqLi4
s0MzN4d3ZtbGuPkJFaWVZ/f4F9fHuAenQlIyJkY2KFhot5eXVMTExgX12JiIdwb22Bf31OTVFBPjtD
QT6Ih4ctKSZ0c26zt9kzKwtUVVldW1qjqrfMzP/R0f+MjKhVU1IxLipZV1Q6Nio0MS6Ji5fExfRnZm
dEQTo8OTa4wuq5uuevu+fR0fAsJBt0eHemuciVlqciHhqZmbR0h5Q1NhxmbXcaGBbExut7dWyZmbk0
NBQ5NjPHyfarq9hzeoeGmaatudE0HAB6t9mzxfJyqck8VWQpFgJ4ort8xe16vuVXgpsYAwCIu9hOXG
ZZaHKHw+aNwN6Gxu46QkZFTlSUvuM0LCMZEw16lqhxmbGJs8qWudVqeolTcoSvwcwmCACYtMeTrsFC
MBWKqbtMPzQMExebsb0WHySVqraIrcO5xdS3w8qntLsLDhAkHRSzt7rL0tfU2NvFzNUXEQnU1evQzf
9COixTRjdmX1dJSEejo9qrq+OhodWrq+WentP///9hJO5aAAAASXRSTlMACI9RQSj8EeTtr/vRovyx
b/71+xXQglk37/71x/7+LAsJGRPld/nWZ3T6TJvcpL3T1v2UJ+3clF43tjNaRuRSrYi96/3DwZ+maH
oAxgAAAAFiS0dE5VhlC78AAAAJcEhZcwAACxIAAAsSAdLdfvwAAAAHdElNRQfkAhgJOyJIb0p3AAAA
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
gXKSofyAL9e3fSU65BheakfgqSPL7toCb+CgL9e/5RZ1ljIyvz9j+cUQkuYGcgUr8IdgG+TvZolWoJ
w4+v+AoLnnYqabnQAAAIRlWElmTU0AKgAAAAgABQESAAMAAAABAAEAAAEaAAUAAAABAAAASgEbAAUA
AAABAAAAUgEoAAMAAAABAAIAAIdpAAQAAAABAAAAWgAAAAAAAABIAAAAAQAAAEgAAAABAAOgAQADAA
AAAQABAACgAgAEAAAAAQAAAQCgAwAEAAAAAQAAALkAAAAAQlvOigAAACV0RVh0ZGF0ZTpjcmVhdGUA
MjAyMC0wMi0yNFQwOTo1NDozMCswMDowMAmDvXAAAAAldEVYdGRhdGU6bW9kaWZ5ADIwMjAtMDItMj
RUMDk6NTQ6MzArMDA6MDB43gXMAAAAEXRFWHRleGlmOkNvbG9yU3BhY2UAMQ+bAkkAAAASdEVYdGV4
aWY6RXhpZk9mZnNldAA5MFmM3psAAAAYdEVYdGV4aWY6UGl4ZWxYRGltZW5zaW9uADI1NtBgr0QAAA
AYdEVYdGV4aWY6UGl4ZWxZRGltZW5zaW9uADE4NWOO35wAAAAASUVORK5CYII=
`
