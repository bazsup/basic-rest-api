<?php

/** @var \Laravel\Lumen\Routing\Router $router */

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
*/

$router->get('/', function () use ($router) {
    return $router->app->version();
});

$router->group(['prefix' => 'posts'], function () use ($router) {
    $router->get('', 'PostController@index');
    $router->post('', 'PostController@create');
    $router->get('/{id:[\d]+}', 'PostController@getById');
    $router->patch('/{id:[\d]+}', 'PostController@update');
    $router->delete('/{id:[\d]+}', 'PostController@delete');
});
