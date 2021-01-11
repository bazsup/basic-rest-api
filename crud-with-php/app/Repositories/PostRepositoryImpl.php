<?php

namespace App\Repositories;

use App\Repositories\PostRepository;
use App\Models\Post;

class PostRepositoryImpl implements PostRepository
{
    protected $post;

    public function __construct(Post $post)
    {
        $this->post = $post;
    }
    
    public function findAll()
    {
        return $this->post->all();
    }
    
    public function findById($id)
    {
        return $this->post->findOrFail($id)->first();
    }

    public function create($data)
    {
        return $this->post->create(array(
            'title' => $data->title,
            'body' => $data->body,
        ));
    }

    public function update($id, $data)
    {
        return tap($this->post->where('id', $id))
            ->update(
                array_filter(
                    array(
                        'title' => $data->title,
                        'body' => $data->body
                    )
                )
            )
            ->first();
    }

    public function destroy($id)
    {
        return $this->post->where('id', $id)->delete();
    }
}