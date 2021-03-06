package tests

import (
	"github.com/jfrogdev/jfrog-cli-go/utils/io/fileutils"
)

const (
	Repo1                      = "jfrog-cli-tests-repo1"
	Repo2                      = "jfrog-cli-tests-repo2"
	Lfs_Repo                   = "jfrog-cli-lfs-repo"
	Out                        = "out"
	DownloadSpec               = "download_spec.json"
	BuildDownloadSpec          = "build_download_spec.json"
	SimpleUploadSpec           = "simple_upload_spec.json"
	UploadEmptyDirs            = "upload_empty_dir_spec.json"
	DownloadEmptyDirs          = "download_empty_dir_spec.json"
	SplittedUploadSpecA        = "splitted_upload_spec_a.json"
	SplittedUploadSpecB        = "splitted_upload_spec_b.json"
	UploadSpec                 = "upload_spec.json"
	DeleteSpec                 = "delete_spec.json"
	DeleteComplexSpec          = "delete_complex_spec.json"
	MoveCopyDeleteSpec         = "move_copy_delete_spec.json"
	PrepareCopy                = "prepare_copy.json"
	Search                     = "search.json"
	SearchRepo2                = "search_repo2.json"
	SearchTxt                  = "search_txt.json"
	SearchMoveDeleteRepoSpec   = "search_move_delete_repo_spec.json"
	CopyByBuildSpec            = "move_copy_delete_by_build_spec.json"
	CpMvDlByBuildAssertSpec    = "copy_by_build_assert_spec.json"
	GitLfsAssertSpec           = "git_lfs_assert_spec.json"
	MoveRepositoryConfig       = "move_repository_config.json"
	SpecsTestRepositoryConfig  = "specs_test_repository_config.json"
	GitLfsTestRepositoryConfig = "git_lfs_test_repository_config.json"
	RepoDetailsUrl             = "api/repositories/"
	CopyItemsSpec             = "copy_items_spec.json"
)

var TxtUploadExpectedRepo1 = []string{
	Repo1 + "/cliTestFile.txt",
}

var SimpleUploadExpectedRepo1 = []string{
	Repo1 + "/flat_recursive/a3.in",
	Repo1 + "/flat_recursive/a1.in",
	Repo1 + "/flat_recursive/a2.in",
	Repo1 + "/flat_recursive/b2.in",
	Repo1 + "/flat_recursive/b3.in",
	Repo1 + "/flat_recursive/b1.in",
	Repo1 + "/flat_recursive/c2.in",
	Repo1 + "/flat_recursive/c1.in",
	Repo1 + "/flat_recursive/c3.in",
}

var SimpleUploadSpecialCharNoRegexExpectedRepo1 = []string{
	Repo1 + "/a1.in",
}

var SingleFileCopy = []string{
	Repo2 + "/path/a1.in",
}

var SingleFileCopyFullPath = []string{
	Repo2 + "/path/inner/a1.in",
}

var SingleInnerFileCopyFullPath = []string{
	Repo2 + "/path/path/inner/a1.in",
}

var FolderCopyTwice = []string{
	Repo2 + "/path/inner/a1.in",
	Repo2 + "/path/path/inner/a1.in",
}

var FolderCopyIntoFolder = []string{
	Repo2 + "/path/path/inner/a1.in",
}

var SingleDirectoryCopyFlat = []string{
	Repo2 + "/inner/a1.in",
}

var AnyItemCopy = []string{
	Repo2 + "/path/inner/a1.in",
	Repo2 + "/someFile",
}

var AnyItemCopyRecursive = []string{
	Repo2 + "/a/b/a1.in",
	Repo2 + "/aFile",
}

var CopyFolderRename = []string{
	Repo2 + "/newPath/inner/a1.in",
}

var AnyItemCopyUsingSpec = []string{
	Repo2 + "/a1.in",
}

var ExplodeUploadExpectedRepo1 = []string{
	Repo1 + "/a/a3.in",
	Repo1 + "/a/a1.in",
	Repo1 + "/a/a2.in",
	Repo1 + "/a/b/b1.in",
	Repo1 + "/a/b/b2.in",
	Repo1 + "/a/b/b3.in",
	Repo1 + "/a/b/c/c1.in",
	Repo1 + "/a/b/c/c2.in",
	Repo1 + "/a/b/c/c3.in",
}

var SimpleUploadExpectedRepo2 = []string{
	Repo2 + "/flat_recursive/a3.in",
	Repo2 + "/flat_recursive/a1.in",
	Repo2 + "/flat_recursive/a2.in",
	Repo2 + "/flat_recursive/b2.in",
	Repo2 + "/flat_recursive/b3.in",
	Repo2 + "/flat_recursive/b1.in",
	Repo2 + "/flat_recursive/c2.in",
	Repo2 + "/flat_recursive/c1.in",
	Repo2 + "/flat_recursive/c3.in",
}

var MassiveMoveExpected = []string{
	Repo2 + "/3_only_flat_recursive_target/a3.in",
	Repo2 + "/3_only_flat_recursive_target/b3.in",
	Repo2 + "/3_only_flat_recursive_target/c3.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a1.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a2.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a3.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b1.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b2.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b3.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c1.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c2.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c3.in",
	Repo2 + "/flat_nonrecursive_target/a1.in",
	Repo2 + "/flat_nonrecursive_target/a2.in",
	Repo2 + "/flat_nonrecursive_target/a3.in",
	Repo2 + "/flat_recursive_target/a1.in",
	Repo2 + "/flat_recursive_target/a2.in",
	Repo2 + "/flat_recursive_target/a3.in",
	Repo2 + "/flat_recursive_target/b1.in",
	Repo2 + "/flat_recursive_target/b2.in",
	Repo2 + "/flat_recursive_target/b3.in",
	Repo2 + "/flat_recursive_target/c1.in",
	Repo2 + "/flat_recursive_target/c2.in",
	Repo2 + "/flat_recursive_target/c3.in",
	Repo2 + "/no_target/a/a1.in",
	Repo2 + "/no_target/a/a2.in",
	Repo2 + "/no_target/a/a3.in",
	Repo2 + "/no_target/a/b/b1.in",
	Repo2 + "/no_target/a/b/b2.in",
	Repo2 + "/no_target/a/b/b3.in",
	Repo2 + "/no_target/a/b/c/c1.in",
	Repo2 + "/no_target/a/b/c/c2.in",
	Repo2 + "/no_target/a/b/c/c3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c3.in",
	Repo2 + "/pattern_placeholder_target/a/a1.in",
	Repo2 + "/pattern_placeholder_target/a/a2.in",
	Repo2 + "/pattern_placeholder_target/a/a3.in",
	Repo2 + "/pattern_placeholder_target/a/b/b1.in",
	Repo2 + "/pattern_placeholder_target/a/b/b2.in",
	Repo2 + "/pattern_placeholder_target/a/b/b3.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c1.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c2.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c3.in",
	Repo2 + "/properties_target/properties_source/a/a1.in",
	Repo2 + "/properties_target/properties_source/a/a2.in",
	Repo2 + "/properties_target/properties_source/a/a3.in",
	Repo2 + "/properties_target/properties_source/a/b/b1.in",
	Repo2 + "/properties_target/properties_source/a/b/b2.in",
	Repo2 + "/properties_target/properties_source/a/b/b3.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c1.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c2.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c3.in",
	Repo2 + "/rename_target/RENAMED.in",
	Repo2 + "/simple_placeholder_target/a/a1.in",
	Repo2 + "/simple_target/a1.in",
	Repo2 + "/flat_nonrecursive_target/b/b1.in",
	Repo2 + "/flat_nonrecursive_target/b/b2.in",
	Repo2 + "/flat_nonrecursive_target/b/b3.in",
	Repo2 + "/flat_nonrecursive_target/b/c/c1.in",
	Repo2 + "/flat_nonrecursive_target/b/c/c2.in",
	Repo2 + "/flat_nonrecursive_target/b/c/c3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c3.in",
}

var BuildCopyExpected = []string{
	Repo1 + "/data/a1.in",
	Repo1 + "/data/a2.in",
	Repo1 + "/data/a3.in",
	Repo1 + "/data/b1.in",
	Repo1 + "/data/b2.in",
	Repo1 + "/data/b3.in",
	Repo2 + "/data/a1.in",
	Repo2 + "/data/a2.in",
	Repo2 + "/data/a3.in",
}

var GitLfsExpected = []string{
	Lfs_Repo + "/objects/4b/f4/4bf4c8c0fef3f5c8cf6f255d1c784377138588c0a9abe57e440bce3ccb350c2e",
}

var BuildMoveExpected = []string{
	Repo1 + "/data/b1.in",
	Repo1 + "/data/b2.in",
	Repo1 + "/data/b3.in",
	Repo2 + "/data/a1.in",
	Repo2 + "/data/a2.in",
	Repo2 + "/data/a3.in",
}

var BuildDeleteExpected = []string{
	Repo1 + "/data/b1.in",
	Repo1 + "/data/b2.in",
	Repo1 + "/data/b3.in",
}

var MassiveDownload = []string{
	Out + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "3_only_flat_recursive" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "3_only_flat_recursive" + fileutils.GetFileSeperator() + "b3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "3_only_flat_recursive" + fileutils.GetFileSeperator() + "c3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "b1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "b2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "b3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "c1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "c2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_flat" + fileutils.GetFileSeperator() + "c3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "defaults_recursive_nonflat" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_nonrecursive" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_nonrecursive" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_nonrecursive" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "b1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "b2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "b3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "c1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "c2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "flat_recursive" + fileutils.GetFileSeperator() + "c3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_nonrecursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_nonrecursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_nonrecursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "nonflat_recursive" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "b3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "properties" + fileutils.GetFileSeperator() + "downloadTestResources" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "b" + fileutils.GetFileSeperator() + "c" + fileutils.GetFileSeperator() + "c3.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "rename" + fileutils.GetFileSeperator() + "a1.out",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "simple" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "simple_placeholder" + fileutils.GetFileSeperator() + "a",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "simple_placeholder" + fileutils.GetFileSeperator() + "a" + fileutils.GetFileSeperator() + "a1.in",
}

var BuildDownload = []string{
	Out,
	Out + fileutils.GetFileSeperator() + "download",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_by_build",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_by_build" + fileutils.GetFileSeperator() + "data",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_by_build" + fileutils.GetFileSeperator() + "data" + fileutils.GetFileSeperator() + "a1.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_by_build" + fileutils.GetFileSeperator() + "data" + fileutils.GetFileSeperator() + "a2.in",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "aql_by_build" + fileutils.GetFileSeperator() + "data" + fileutils.GetFileSeperator() + "a3.in",
}

var BuildSimpleDownload = []string{
	Out,
	Out + fileutils.GetFileSeperator() + "download",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "simple_by_build",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "simple_by_build" + fileutils.GetFileSeperator() + "data",
	Out + fileutils.GetFileSeperator() + "download" + fileutils.GetFileSeperator() + "simple_by_build" + fileutils.GetFileSeperator() + "data" + fileutils.GetFileSeperator() + "b1.in",
}

var MassiveUpload = []string{
	Repo1 + "/spec-copy-test/3_only_flat_recursive/a3.in",
	Repo1 + "/spec-copy-test/3_only_flat_recursive/b3.in",
	Repo1 + "/spec-copy-test/3_only_flat_recursive/c3.in",
	Repo1 + "/spec-copy-test/copy-to-existing/a1.in",
	Repo1 + "/spec-copy-test/copy-to-existing/a2.in",
	Repo1 + "/spec-copy-test/copy-to-existing/a3.in",
	Repo1 + "/spec-copy-test/copy-to-existing/b1.in",
	Repo1 + "/spec-copy-test/copy-to-existing/b2.in",
	Repo1 + "/spec-copy-test/copy-to-existing/b3.in",
	Repo1 + "/spec-copy-test/copy-to-existing/c1.in",
	Repo1 + "/spec-copy-test/copy-to-existing/c2.in",
	Repo1 + "/spec-copy-test/copy-to-existing/c3.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/a1.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/a2.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/a3.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/b1.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/b2.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/b3.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/c1.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/c2.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonflat/c3.in",
	Repo1 + "/spec-copy-test/flat_nonrecursive/a1.in",
	Repo1 + "/spec-copy-test/flat_nonrecursive/a2.in",
	Repo1 + "/spec-copy-test/flat_nonrecursive/a3.in",
	Repo1 + "/spec-copy-test/flat_recursive/a1.in",
	Repo1 + "/spec-copy-test/flat_recursive/a2.in",
	Repo1 + "/spec-copy-test/flat_recursive/a3.in",
	Repo1 + "/spec-copy-test/flat_recursive/b1.in",
	Repo1 + "/spec-copy-test/flat_recursive/b2.in",
	Repo1 + "/spec-copy-test/flat_recursive/b3.in",
	Repo1 + "/spec-copy-test/flat_recursive/c1.in",
	Repo1 + "/spec-copy-test/flat_recursive/c2.in",
	Repo1 + "/spec-copy-test/flat_recursive/c3.in",
	Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b1.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b2.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b3.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c1.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c2.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a+a/a1.in",
	Repo1 + "/spec-copy-test/simple/a1.in",
}

var PropsExpected = []string{
	Repo1 + "/spec-copy-test/properties/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a+a/a1.in",
}

var Delete1 = []string{
	Repo2 + "/3_only_flat_recursive_target/a3.in",
	Repo2 + "/3_only_flat_recursive_target/b3.in",
	Repo2 + "/3_only_flat_recursive_target/c3.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a1.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a2.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/a3.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b1.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b2.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/b3.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c1.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c2.in",
	Repo2 + "/defaults_recursive_nonflat_target/defaults_recursive_nonflat_source/a/b/c/c3.in",
	Repo2 + "/flat_nonrecursive_target/a1.in",
	Repo2 + "/flat_nonrecursive_target/a2.in",
	Repo2 + "/flat_nonrecursive_target/a3.in",
	Repo2 + "/flat_recursive_target/a1.in",
	Repo2 + "/flat_recursive_target/a2.in",
	Repo2 + "/flat_recursive_target/a3.in",
	Repo2 + "/flat_recursive_target/b1.in",
	Repo2 + "/flat_recursive_target/b2.in",
	Repo2 + "/flat_recursive_target/b3.in",
	Repo2 + "/flat_recursive_target/c1.in",
	Repo2 + "/flat_recursive_target/c2.in",
	Repo2 + "/flat_recursive_target/c3.in",
	Repo2 + "/no_target/a/a1.in",
	Repo2 + "/no_target/a/a2.in",
	Repo2 + "/no_target/a/a3.in",
	Repo2 + "/no_target/a/b/b1.in",
	Repo2 + "/no_target/a/b/b2.in",
	Repo2 + "/no_target/a/b/b3.in",
	Repo2 + "/no_target/a/b/c/c1.in",
	Repo2 + "/no_target/a/b/c/c2.in",
	Repo2 + "/no_target/a/b/c/c3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a3.in",
	Repo2 + "/pattern_placeholder_target/a/a1.in",
	Repo2 + "/pattern_placeholder_target/a/a2.in",
	Repo2 + "/pattern_placeholder_target/a/a3.in",
	Repo2 + "/pattern_placeholder_target/a/b/b1.in",
	Repo2 + "/pattern_placeholder_target/a/b/b2.in",
	Repo2 + "/pattern_placeholder_target/a/b/b3.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c1.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c2.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c3.in",
	Repo2 + "/properties_target/properties_source/a/a1.in",
	Repo2 + "/properties_target/properties_source/a/a2.in",
	Repo2 + "/properties_target/properties_source/a/a3.in",
	Repo2 + "/properties_target/properties_source/a/b/b1.in",
	Repo2 + "/properties_target/properties_source/a/b/b2.in",
	Repo2 + "/properties_target/properties_source/a/b/b3.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c1.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c2.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c3.in",
	Repo2 + "/rename_target/RENAMED.in",
	Repo2 + "/simple_placeholder_target/a/a1.in",
	Repo2 + "/simple_target/a1.in",
	Repo2 + "/flat_nonrecursive_target/b/b1.in",
	Repo2 + "/flat_nonrecursive_target/b/b2.in",
	Repo2 + "/flat_nonrecursive_target/b/b3.in",
	Repo2 + "/flat_nonrecursive_target/b/c/c1.in",
	Repo2 + "/flat_nonrecursive_target/b/c/c2.in",
	Repo2 + "/flat_nonrecursive_target/b/c/c3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/b3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/b/c/c3.in",
}

var DeleteDisplyedFiles = []string{
	Repo2 + "/3_only_flat_recursive_source/a/b/",
	Repo2 + "/3_only_flat_recursive_source/a/a1.in",
	Repo2 + "/3_only_flat_recursive_source/a/a2.in",
	Repo2 + "/3_only_flat_recursive_source/a/a3.in",
	Repo2 + "/flat_recursive_source/a/b/c/",
	Repo2 + "/flat_recursive_source/a/b/b1.in",
	Repo2 + "/flat_recursive_source/a/b/b2.in",
	Repo2 + "/flat_recursive_source/a/b/b3.in",
	Repo2 + "/defaults_recursive_nonflat_source/a/a1.in",
	Repo2 + "/defaults_recursive_nonflat_source/a/a2.in",
	Repo2 + "/defaults_recursive_nonflat_source/a/a3.in",
	Repo2 + "/defaults_recursive_nonflat_source/a/b/",
	Repo2 + "/flat_nonrecursive_source/a/b/c/",
	Repo2 + "/flat_nonrecursive_source/a/b/b1.in",
	Repo2 + "/flat_nonrecursive_source/a/b/b2.in",
	Repo2 + "/flat_nonrecursive_source/a/b/b3.in",
}