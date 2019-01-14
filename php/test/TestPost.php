<?php
namespace test;

use GuzzleHttp\Pool;
use GuzzleHttp\Client;
use GuzzleHttp\ClientInterface;
use GuzzleHttp\Exception\RequestException;
use Psr\Http\Message\ResponseInterface;

require __DIR__ . '/TestPost.php';

class TestPost {

    public function run()
    {
        $client = new Client([  
            // You can set any number of default request options.
            'timeout'  => 10,
            'debug' => true,
            'decode_content' => true,
            'headers' => [
                'content-type' => 'application/json; charset=UTF-8',
            ]
        ]);
        
        
        $url = 'https://api.ifcert.org.cn/p2p/userInfo/test';
        echo $url . PHP_EOL;
        $resp = $client->post($url,[
            'form_params' => [
                'apiKey' => '64a0f148a08dd6febf431a7cbb69832f84d762f22d1f86e195887c61f71f3af1',
                'msg' => '{"dataList":[{"version":"1.5","userCreateTime":"2017-05-17 00:00:00","sourceCode":"CERT20181122006","userStatus":"2","userType":"1","userAttr":"1","userName":"张三","countries":"1","cardType":"1","userIdcard":"147852dfghjklfghjkl","userIdcardHash":"a0a1cb15053a1dfd98a0fb167da9b417d4d65f636595db31904313605264068a","userPhone":"18330******","userPhoneHash":"$2a$10$zcPH1/imspdxSsCnyLOEYO.CGBz3W1RyXT4g5AmB23HZDVHDUuRdW","userUuid":"$2a$10$bhLUTsbY2qD5TCGYzrEEp.","userLawperson":"-1","userFund":"-1","userProvince":"-1","userAddress":"-1","registerDate":"-1","userMail":"qq@163.com","userAscription":"140830","userAge":"19","userSex":"1","phoneAscription":"140100","riskRating":"A","userList":[{"userPay":"汇付天下","userPayAccount":"HFTX00001","userBank":"汇付天下北京分部","userBankAccount":"60100000001"},{"userPay":"如意支付","userPayAccount":"RYF00001","userBank":"北京如意支付","userBankAccount":"60100000001"}]}],"version":"1.5","batchNum":"CERT20181122006_20181212_1544620728503","checkCode":"0dd52c64806849c29a30184aa2186583","totalNum":"1","sentTime":"2018-12-12 21:18:48","sourceCode":"CERT20181122006","infType":"1","dataType":"1","timestamp":"1544620728503","nonce":"aeb855db5cae447ea75f6f5be8e716af"}',
            ],
            'headers' => [
                'content-type' => 'application/x-www-form-urlencoded;charset=UTF-8',
            ]
        ]);
        echo $resp->getBody();       
        echo PHP_EOL;     
    }
}
