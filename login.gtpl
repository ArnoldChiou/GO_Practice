<html>
<head>
<title></title>
</head>
<body>
<form action="/login?username=astaxie" method="post">
    <input type="checkbox" name="interest" value="football">足球
    <input type="checkbox" name="interest" value="basketball">籃球
    <input type="checkbox" name="interest" value="tennis">網球
    使用者名稱:<input type="text" name="username">
    密碼:<input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登入">
</form>
<script>
alert("hello")
</script>
</body>
</html>