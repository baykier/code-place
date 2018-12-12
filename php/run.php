<?php

require __DIR__ . '/vendor/autoload.php';

use test\TestGet;
use test\TestPost;

// $get = new TestGet();
// $get->run();

$batch = new TestPost;
$batch->run();