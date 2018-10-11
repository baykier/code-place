<?php

$n = -1;

printf("%d bit code is %08b\n",$n, $n);

$m = unpack("H*",'a');
var_dump($m);
$s = 0x12;
var_dump($s);
var_dump(012);//8 进制