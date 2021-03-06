// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codecommitiface provides an interface to enable mocking the AWS CodeCommit service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codecommitiface

import (
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
)

// CodeCommitAPI provides an interface to enable mocking the
// codecommit.CodeCommit service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeCommit.
//    func myFunc(svc codecommitiface.CodeCommitAPI) bool {
//        // Make svc.BatchGetRepositories request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := codecommit.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeCommitClient struct {
//        codecommitiface.CodeCommitAPI
//    }
//    func (m *mockCodeCommitClient) BatchGetRepositories(input *codecommit.BatchGetRepositoriesInput) (*codecommit.BatchGetRepositoriesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeCommitClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CodeCommitAPI interface {
	BatchGetRepositoriesRequest(*codecommit.BatchGetRepositoriesInput) codecommit.BatchGetRepositoriesRequest

	CreateBranchRequest(*codecommit.CreateBranchInput) codecommit.CreateBranchRequest

	CreatePullRequestRequest(*codecommit.CreatePullRequestInput) codecommit.CreatePullRequestRequest

	CreateRepositoryRequest(*codecommit.CreateRepositoryInput) codecommit.CreateRepositoryRequest

	DeleteBranchRequest(*codecommit.DeleteBranchInput) codecommit.DeleteBranchRequest

	DeleteCommentContentRequest(*codecommit.DeleteCommentContentInput) codecommit.DeleteCommentContentRequest

	DeleteFileRequest(*codecommit.DeleteFileInput) codecommit.DeleteFileRequest

	DeleteRepositoryRequest(*codecommit.DeleteRepositoryInput) codecommit.DeleteRepositoryRequest

	DescribePullRequestEventsRequest(*codecommit.DescribePullRequestEventsInput) codecommit.DescribePullRequestEventsRequest

	GetBlobRequest(*codecommit.GetBlobInput) codecommit.GetBlobRequest

	GetBranchRequest(*codecommit.GetBranchInput) codecommit.GetBranchRequest

	GetCommentRequest(*codecommit.GetCommentInput) codecommit.GetCommentRequest

	GetCommentsForComparedCommitRequest(*codecommit.GetCommentsForComparedCommitInput) codecommit.GetCommentsForComparedCommitRequest

	GetCommentsForPullRequestRequest(*codecommit.GetCommentsForPullRequestInput) codecommit.GetCommentsForPullRequestRequest

	GetCommitRequest(*codecommit.GetCommitInput) codecommit.GetCommitRequest

	GetDifferencesRequest(*codecommit.GetDifferencesInput) codecommit.GetDifferencesRequest

	GetFileRequest(*codecommit.GetFileInput) codecommit.GetFileRequest

	GetFolderRequest(*codecommit.GetFolderInput) codecommit.GetFolderRequest

	GetMergeConflictsRequest(*codecommit.GetMergeConflictsInput) codecommit.GetMergeConflictsRequest

	GetPullRequestRequest(*codecommit.GetPullRequestInput) codecommit.GetPullRequestRequest

	GetRepositoryRequest(*codecommit.GetRepositoryInput) codecommit.GetRepositoryRequest

	GetRepositoryTriggersRequest(*codecommit.GetRepositoryTriggersInput) codecommit.GetRepositoryTriggersRequest

	ListBranchesRequest(*codecommit.ListBranchesInput) codecommit.ListBranchesRequest

	ListPullRequestsRequest(*codecommit.ListPullRequestsInput) codecommit.ListPullRequestsRequest

	ListRepositoriesRequest(*codecommit.ListRepositoriesInput) codecommit.ListRepositoriesRequest

	MergePullRequestByFastForwardRequest(*codecommit.MergePullRequestByFastForwardInput) codecommit.MergePullRequestByFastForwardRequest

	PostCommentForComparedCommitRequest(*codecommit.PostCommentForComparedCommitInput) codecommit.PostCommentForComparedCommitRequest

	PostCommentForPullRequestRequest(*codecommit.PostCommentForPullRequestInput) codecommit.PostCommentForPullRequestRequest

	PostCommentReplyRequest(*codecommit.PostCommentReplyInput) codecommit.PostCommentReplyRequest

	PutFileRequest(*codecommit.PutFileInput) codecommit.PutFileRequest

	PutRepositoryTriggersRequest(*codecommit.PutRepositoryTriggersInput) codecommit.PutRepositoryTriggersRequest

	TestRepositoryTriggersRequest(*codecommit.TestRepositoryTriggersInput) codecommit.TestRepositoryTriggersRequest

	UpdateCommentRequest(*codecommit.UpdateCommentInput) codecommit.UpdateCommentRequest

	UpdateDefaultBranchRequest(*codecommit.UpdateDefaultBranchInput) codecommit.UpdateDefaultBranchRequest

	UpdatePullRequestDescriptionRequest(*codecommit.UpdatePullRequestDescriptionInput) codecommit.UpdatePullRequestDescriptionRequest

	UpdatePullRequestStatusRequest(*codecommit.UpdatePullRequestStatusInput) codecommit.UpdatePullRequestStatusRequest

	UpdatePullRequestTitleRequest(*codecommit.UpdatePullRequestTitleInput) codecommit.UpdatePullRequestTitleRequest

	UpdateRepositoryDescriptionRequest(*codecommit.UpdateRepositoryDescriptionInput) codecommit.UpdateRepositoryDescriptionRequest

	UpdateRepositoryNameRequest(*codecommit.UpdateRepositoryNameInput) codecommit.UpdateRepositoryNameRequest
}

var _ CodeCommitAPI = (*codecommit.CodeCommit)(nil)
