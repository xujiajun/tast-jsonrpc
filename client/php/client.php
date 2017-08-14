<?php
require "./JsonRPC.php";
require "./Runtime.php";

$runtime= new Runtime();
$runtime->start();
$client = new JsonRPC("127.0.0.1", 1234);

for ($i=0; $i < 10000; $i++) { 
    $r = $client->Call("MyMath.Add",array('num1'=>1,'num2'=>2));
}

$runtime->stop();
$QPS = 10000/($runtime->spent()/1000);
echo "process finished! 执行时间: ".$runtime->spent()." 毫秒. QPS: ($QPS)".PHP_EOL;