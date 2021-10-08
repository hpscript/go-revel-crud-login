# About
This is a CRUD sample of Revel, the framework of Go. We use mysql for the database.

### Create DB table and data
```
use test;
create table baseballs(
	id int primary key auto_increment,
	name varchar(255),
	manager varchar(255),
	home varchar(255)
);
```

```
insert into baseballs(name, manager, home) values ('東京ヤクルトスワローズ', '高津臣吾', '明治神宮野球場');
insert into baseballs(name, manager, home) values ('阪神タイガース', '矢野燿大', '阪神甲子園球場');
insert into baseballs(name, manager, home) values ('読売ジャイアンツ', '原 辰徳', '東京ドーム');
insert into baseballs(name, manager, home) values ('広島東洋カープ', '佐々岡真司', 'MAZDA Zoom-Zoom スタジアム 広島');
insert into baseballs(name, manager, home) values ('中日ドラゴンズ', '与田剛', 'バンテリンドーム ナゴヤ');
insert into baseballs(name, manager, home) values ('横浜DeNAベイスターズ', '三浦大輔', '横浜スタジアム');
insert into baseballs(name, manager, home) values ('福岡ソフトバンクホークス', '工藤公康', '福岡PayPayドーム');
insert into baseballs(name, manager, home) values ('千葉ロッテマリーンズ', '井口資仁', 'ZOZOマリンスタジアム');
insert into baseballs(name, manager, home) values ('オリックス・バファローズ', '中嶋聡', '大阪市神戸市');
insert into baseballs(name, manager, home) values ('北海道日本ハムファイターズ', '栗山英樹', '札幌ドーム');
insert into baseballs(name, manager, home) values ('埼玉西武ライオンズ', '辻発彦', 'メットライフドーム');
insert into baseballs(name, manager, home) values ('東北楽天ゴールデンイーグルス', '石井一久', '楽天生命パーク宮城');
```

### Run
$ revel run -a app
