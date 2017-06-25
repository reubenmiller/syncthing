// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package fs

import "time"

type logFilesystem struct {
	Filesystem
}

func (fs *logFilesystem) Chmod(name string, mode FileMode) error {
	err := fs.Filesystem.Chmod(name, mode)
	l.Debugln(fs.Type(), fs.URI(), "Chmod", name, mode, err)
	return err
}

func (fs *logFilesystem) Chtimes(name string, atime time.Time, mtime time.Time) error {
	err := fs.Filesystem.Chtimes(name, atime, mtime)
	l.Debugln(fs.Type(), fs.URI(), "Chtimes", name, atime, mtime, err)
	return err
}

func (fs *logFilesystem) Create(name string) (File, error) {
	file, err := fs.Filesystem.Create(name)
	l.Debugln(fs.Type(), fs.URI(), "Create", name, file, err)
	return file, err
}

func (fs *logFilesystem) CreateSymlink(name, target string) error {
	err := fs.Filesystem.CreateSymlink(name, target)
	l.Debugln(fs.Type(), fs.URI(), "CreateSymlink", name, target, err)
	return err
}

func (fs *logFilesystem) DirNames(name string) ([]string, error) {
	names, err := fs.Filesystem.DirNames(name)
	l.Debugln(fs.Type(), fs.URI(), "DirNames", name, names, err)
	return names, err
}

func (fs *logFilesystem) Lstat(name string) (FileInfo, error) {
	info, err := fs.Filesystem.Lstat(name)
	l.Debugln(fs.Type(), fs.URI(), "Lstat", name, info, err)
	return info, err
}

func (fs *logFilesystem) Mkdir(name string, perm FileMode) error {
	err := fs.Filesystem.Mkdir(name, perm)
	l.Debugln(fs.Type(), fs.URI(), "Mkdir", name, perm, err)
	return err
}

func (fs *logFilesystem) MkdirAll(name string, perm FileMode) error {
	err := fs.Filesystem.MkdirAll(name, perm)
	l.Debugln(fs.Type(), fs.URI(), "MkdirAll", name, perm, err)
	return err
}

func (fs *logFilesystem) Open(name string) (File, error) {
	file, err := fs.Filesystem.Open(name)
	l.Debugln(fs.Type(), fs.URI(), "Open", name, file, err)
	return file, err
}

func (fs *logFilesystem) OpenFile(name string, flags int, mode FileMode) (File, error) {
	file, err := fs.Filesystem.OpenFile(name, flags, mode)
	l.Debugln(fs.Type(), fs.URI(), "OpenFile", name, flags, mode, file, err)
	return file, err
}

func (fs *logFilesystem) ReadSymlink(name string) (string, error) {
	target, err := fs.Filesystem.ReadSymlink(name)
	l.Debugln(fs.Type(), fs.URI(), "ReadSymlink", name, target, err)
	return target, err
}

func (fs *logFilesystem) Remove(name string) error {
	err := fs.Filesystem.Remove(name)
	l.Debugln(fs.Type(), fs.URI(), "Remove", name, err)
	return err
}

func (fs *logFilesystem) RemoveAll(name string) error {
	err := fs.Filesystem.RemoveAll(name)
	l.Debugln(fs.Type(), fs.URI(), "RemoveAll", name, err)
	return err
}

func (fs *logFilesystem) Rename(oldname, newname string) error {
	err := fs.Filesystem.Rename(oldname, newname)
	l.Debugln(fs.Type(), fs.URI(), "Rename", oldname, newname, err)
	return err
}

func (fs *logFilesystem) Stat(name string) (FileInfo, error) {
	info, err := fs.Filesystem.Stat(name)
	l.Debugln(fs.Type(), fs.URI(), "Stat", name, info, err)
	return info, err
}

func (fs *logFilesystem) SymlinksSupported() bool {
	supported := fs.Filesystem.SymlinksSupported()
	l.Debugln(fs.Type(), fs.URI(), "SymlinksSupported", supported)
	return supported
}

func (fs *logFilesystem) Walk(root string, walkFn WalkFunc) error {
	err := fs.Filesystem.Walk(root, walkFn)
	l.Debugln(fs.Type(), fs.URI(), "Walk", root, walkFn, err)
	return err
}

func (fs *logFilesystem) Show(name string) error {
	err := fs.Filesystem.Show(name)
	l.Debugln(fs.Type(), fs.URI(), "Show", name, err)
	return err
}

func (fs *logFilesystem) Hide(name string) error {
	err := fs.Filesystem.Hide(name)
	l.Debugln(fs.Type(), fs.URI(), "Hide", name, err)
	return err
}

func (fs *logFilesystem) Glob(name string) ([]string, error) {
	names, err := fs.Filesystem.Glob(name)
	l.Debugln(fs.Type(), fs.URI(), "Glob", name, names, err)
	return names, err
}

func (fs *logFilesystem) SyncDir(name string) error {
	err := fs.Filesystem.SyncDir(name)
	l.Debugln(fs.Type(), fs.URI(), "SyncDir", name, err)
	return err
}

func (fs *logFilesystem) Roots() ([]string, error) {
	roots, err := fs.Filesystem.Roots()
	l.Debugln(fs.Type(), fs.URI(), "Roots", roots, err)
	return roots, err
}

func (fs *logFilesystem) Usage(name string) (Usage, error) {
	usage, err := fs.Filesystem.Usage(name)
	l.Debugln(fs.Type(), fs.URI(), "Usage", name, usage, err)
	return usage, err
}