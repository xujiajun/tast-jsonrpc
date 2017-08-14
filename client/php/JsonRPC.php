<?php

class JsonRPC
{
    private $conn;
    private $id;
    private $host;
    private $port;

    public function __construct($host, $port)
    {
        $this->host = $host;
        $this->port = $port;
        $this->id = 1;
    }

    private function Dial()
    {
        $this->conn = fsockopen($this->host, $this->port, $errno, $errstr, 3);

        if (!$this->conn) {
            return "JsonRPC Dial Failed: $errstr ($errno)";
        }

        stream_set_timeout($this->conn, 3);
        $info = stream_get_meta_data($this->conn);
        if ($info['timed_out']) {
            fclose($this->conn);
            return "JsonRPC Init Time Out";
        }

        return NULL;
    }

    public function Call($method, $params)
    {
        if (!$this->conn) {
            $dialResult = $this->Dial();
            if ($dialResult !== NULL) {
                return $dialResult;
            }
        }
        $err = fwrite($this->conn, json_encode(array(
                'method' => $method,
                'params' => array($params),
                'id' => $this->id++,
            )) . PHP_EOL);
        if ($err === false) {
            return false;
        }

        $line = fgets($this->conn);
        if ($line === false) {
            return NULL;
        }
        return json_decode($line, true);
    }
}