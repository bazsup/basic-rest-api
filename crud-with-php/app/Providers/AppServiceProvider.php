<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;

use App\Repositories\PostRepository;
use App\Repositories\PostRepositoryImpl;

class AppServiceProvider extends ServiceProvider
{

    public function boot() {}
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {        
        $this->app->bind('App\Repositories\\PostRepository', 'App\Repositories\\PostRepositoryImpl');
    }
}
