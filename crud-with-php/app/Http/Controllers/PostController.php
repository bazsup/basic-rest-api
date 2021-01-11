<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Repositories\PostRepository;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;

class PostController extends Controller
{
    protected $post;

    public function __construct(PostRepository $post)
    {
        $this->post = $post;
    }

    public function index()
    {
        return $this->post->findAll();
    }

    public function getById($id)
    {
        $result = $this->post->findById($id);
        return response()->json($result);
    }

    public function create(Request $request)
    {
        $this->validate($request, [
            'title' => 'required',
            'body' => 'required',
        ]);

        $res = $this->post->create($request);
        return response()->json($res, 201);
    }

    public function update($id, Request $request)
    {
        $result = $this->post->update($id, $request);
        if ($result == null)
        {
            throw new NotFoundHttpException();
        }

        return response()->json($result);
    }

    public function delete($id)
    {
        $result = $this->post->destroy($id);
        if ($result == 0)
        {
            throw new NotFoundHttpException();
        }

        return response()->json(null, 204);
    }
}