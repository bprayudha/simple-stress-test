<?php
namespace App\Http\Controllers;

use App\Http\Controllers\Controller;

class HelloController extends Controller
{
    public function index()
    {
        return response()->json(["hello" => "world"]);
    }
}
