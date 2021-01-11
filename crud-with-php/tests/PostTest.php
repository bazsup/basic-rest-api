<?php

use Laravel\Lumen\Testing\DatabaseMigrations;
use App\Models\Post;
use Faker\Factory;

class PostTest extends TestCase
{
    use DatabaseMigrations;

    private $faker;

    public function setUp(): void
    {
        parent::setUp();
        $this->faker = Factory::create();
    }
    
    public function testGetAllPosts() {
        $post = Post::create(array(
            'title' => $this->faker->sentence(),
            'body' => $this->faker->paragraph(),
        ));
        $this->json('GET', '/posts')
            ->seeStatusCode(200)
            ->seeJsonEquals([
                $post
            ]);
    }

    public function testCreatePost() {
        $expected = ['title' => 'title 1', 'body' => 'body 1'];

        $this->json('POST', '/posts', $expected)
            ->seeStatusCode(201)
            ->seeJsonContains($expected);
    }

    public function testGetPostById_Success() {
        $expected = ['title' => 'title "get by id"', 'body' => 'body 1'];
        Post::create($expected);

        $this->json('GET', '/posts/1')
            ->seeStatusCode(200)
            ->seeJsonContains($expected);
    }

    public function testGetPostById_NotFound() {
        $this->json('GET', '/posts/2')
            ->seeStatusCode(404)
            ->seeJsonEquals(['message' => 'Not Found']);
    }

    public function testupdatePost() {
        $expected = ['title' => 'title "updated post"'];
        Post::create(['title' => 'title 1', 'body' => 'body 1']);
        
        $this->json('PATCH', '/posts/1', $expected)
        ->seeStatusCode(200)
        ->seeJsonContains($expected);
    }
    
    public function testupdatePost_NotFound() {
        $expected = ['title' => 'title "updated post"'];
        
        $this->json('PATCH', '/posts/1', $expected)
        ->seeStatusCode(404)
        ->seeJsonEquals(['message' => 'Not Found']);
    }
    
    public function testDeletePost() {
        Post::create(['title' => 'title 1', 'body' => 'body 1']);

        $this->json('DELETE', '/posts/1')
            ->seeStatusCode(204)
            ->notSeeInDatabase('posts', ['title' => 'title 1']);
    }

    public function testDeletePost_NotFound() {
        $this->json('DELETE', '/posts/1')
            ->seeStatusCode(404)
            ->seeJsonEquals(['message' => 'Not Found']);
    }
}