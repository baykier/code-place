<?php
namespace test;

use GuzzleHttp\Pool;
use GuzzleHttp\Client;
use GuzzleHttp\ClientInterface;
use GuzzleHttp\Exception\RequestException;
use Psr\Http\Message\ResponseInterface;

class TestGet {

    public function run()
    {
        $client = new Client([  
            // You can set any number of default request options.
            'timeout'  => 5.0,
        ]);
        
        
        $url = 'http://baidu.com';
        echo $url . PHP_EOL;
        $promise = $client->getAsync($url);
        $promise->then(
            function (ResponseInterface $res) {
                echo $res->getStatusCode() . "\n";
            },
            function (RequestException $e) {
                echo $e->getMessage() . "\n";
                echo $e->getRequest()->getMethod();
            }
        )->wait();
        echo 'this is the first';
        echo PHP_EOL;
        //$res = $promise->wait();
        //echo $res->getBody();
    }
}
