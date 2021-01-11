<?php

namespace App\Repositories;

interface PostRepository {
    public function findAll();
    public function findById($id);
    public function create($data);
    public function update($id, $data);
    public function destroy($id);
}