/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package identity_provider

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"k8s.io/klog/klogr"

	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services/eks/mock_eksiface"
)

func TestEKSAddonPlan(t *testing.T) {
	clusterName := "default.cluster"
	log := klogr.New()

	testCases := []struct {
		name                    string
		currentIdentityProvider *OidcIdentityProviderConfig
		desiredIdentityProvider *OidcIdentityProviderConfig
		expect                  func(m *mock_eksiface.MockEKSAPIMockRecorder)
		expectCreateError       bool
		expectDoError           bool
	}{
		{
			name: "no desired and no installed",
			expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
				// Do nothing
			},
			expectCreateError: false,
			expectDoError:     false,
		},
		//{
		//	name: "no installed and 1 desired",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.
		//			CreateAddon(gomock.Eq(&eks.CreateAddonInput{
		//				AddonName:        aws.String(addon1Name),
		//				AddonVersion:     aws.String(addon1version),
		//				ClusterName:      aws.String(clusterName),
		//				ResolveConflicts: aws.String(eks.ResolveConflictsOverwrite),
		//				Tags:             convertTags(createTags()),
		//			})).
		//			Return(&eks.CreateAddonOutput{
		//				Addon: &eks.Addon{
		//					AddonArn:     aws.String(addonARN),
		//					AddonName:    aws.String(addon1Name),
		//					AddonVersion: aws.String(addon1version),
		//					ClusterName:  aws.String(clusterName),
		//					CreatedAt:    &created,
		//					ModifiedAt:   &created,
		//					Status:       aws.String(addonStatusCreating),
		//					Tags:         convertTags(createTags()),
		//				},
		//			}, nil)
		//
		//		m.WaitUntilAddonActive(gomock.Eq(&eks.DescribeAddonInput{
		//			AddonName:   aws.String(addon1Name),
		//			ClusterName: aws.String(clusterName),
		//		})).Return(nil)
		//	},
		//	desiredAddons: []*EKSAddon{
		//		createDesiredAddon(addon1Name, addon1version),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 1 desired - both same and installed active",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		// No Action expected
		//	},
		//	desiredAddons: []*EKSAddon{
		//		createDesiredAddon(addon1Name, addon1version),
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1version, addonARN, addonStatusActive),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 1 desired - both same and installed is creating",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.WaitUntilAddonActive(gomock.Eq(&eks.DescribeAddonInput{
		//			AddonName:   aws.String(addon1Name),
		//			ClusterName: aws.String(clusterName),
		//		})).Return(nil)
		//	},
		//	desiredAddons: []*EKSAddon{
		//		createDesiredAddon(addon1Name, addon1version),
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1version, addonARN, addonStatusCreating),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 1 desired - version upgrade",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.
		//			UpdateAddon(gomock.Eq(&eks.UpdateAddonInput{
		//				AddonName:        aws.String(addon1Name),
		//				AddonVersion:     aws.String(addon1Upgrade),
		//				ClusterName:      aws.String(clusterName),
		//				ResolveConflicts: aws.String(eks.ResolveConflictsOverwrite),
		//			})).
		//			Return(&eks.UpdateAddonOutput{
		//				Update: &eks.Update{
		//					CreatedAt: &created,
		//					Id:        aws.String("someid"),
		//					Status:    aws.String(addonStatusUpdating),
		//					Type:      aws.String(eks.UpdateTypeVersionUpdate),
		//				},
		//			}, nil)
		//		m.WaitUntilAddonActive(gomock.Eq(&eks.DescribeAddonInput{
		//			AddonName:   aws.String(addon1Name),
		//			ClusterName: aws.String(clusterName),
		//		})).Return(nil)
		//	},
		//	desiredAddons: []*EKSAddon{
		//		createDesiredAddon(addon1Name, addon1Upgrade),
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1version, addonARN, addonStatusActive),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 1 desired - version upgrade in progress",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.WaitUntilAddonActive(gomock.Eq(&eks.DescribeAddonInput{
		//			AddonName:   aws.String(addon1Name),
		//			ClusterName: aws.String(clusterName),
		//		})).Return(nil)
		//	},
		//	desiredAddons: []*EKSAddon{
		//		createDesiredAddon(addon1Name, addon1Upgrade),
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1Upgrade, addonARN, addonStatusUpdating),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 1 desired - tags upgrade",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.
		//			TagResource(gomock.Eq(&eks.TagResourceInput{
		//				ResourceArn: &addonARN,
		//				Tags:        convertTags(createTagsAdditional()),
		//			})).
		//			Return(&eks.TagResourceOutput{}, nil)
		//	},
		//	desiredAddons: []*EKSAddon{
		//		createDesiredAddonExtraTag(addon1Name, addon1version),
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1version, addonARN, addonStatusActive),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 1 desired - version & tags upgrade",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.
		//			TagResource(gomock.Eq(&eks.TagResourceInput{
		//				ResourceArn: &addonARN,
		//				Tags:        convertTags(createTagsAdditional()),
		//			})).
		//			Return(&eks.TagResourceOutput{}, nil)
		//		m.
		//			UpdateAddon(gomock.Eq(&eks.UpdateAddonInput{
		//				AddonName:        aws.String(addon1Name),
		//				AddonVersion:     aws.String(addon1Upgrade),
		//				ClusterName:      aws.String(clusterName),
		//				ResolveConflicts: aws.String(eks.ResolveConflictsOverwrite),
		//			})).
		//			Return(&eks.UpdateAddonOutput{
		//				Update: &eks.Update{
		//					CreatedAt: &created,
		//					Id:        aws.String("someid"),
		//					Status:    aws.String(addonStatusUpdating),
		//					Type:      aws.String(eks.UpdateTypeVersionUpdate),
		//				},
		//			}, nil)
		//
		//		m.WaitUntilAddonActive(gomock.Eq(&eks.DescribeAddonInput{
		//			AddonName:   aws.String(addon1Name),
		//			ClusterName: aws.String(clusterName),
		//		})).Return(nil)
		//	},
		//	desiredAddons: []*EKSAddon{
		//		createDesiredAddonExtraTag(addon1Name, addon1Upgrade),
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1version, addonARN, addonStatusActive),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 0 desired - delete addon",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.
		//			DeleteAddon(gomock.Eq(&eks.DeleteAddonInput{
		//				AddonName:   &addon1Name,
		//				ClusterName: &clusterName,
		//			})).
		//			Return(&eks.DeleteAddonOutput{
		//				Addon: &eks.Addon{
		//					AddonArn:     aws.String(addonARN),
		//					AddonName:    aws.String(addon1Name),
		//					AddonVersion: aws.String(addon1version),
		//					ClusterName:  aws.String(clusterName),
		//					CreatedAt:    &created,
		//					ModifiedAt:   &created,
		//					Status:       aws.String(addonStatusDeleting),
		//					Tags:         convertTags(createTags()),
		//				},
		//			}, nil)
		//
		//		m.WaitUntilAddonDeleted(gomock.Eq(&eks.DescribeAddonInput{
		//			AddonName:   aws.String(addon1Name),
		//			ClusterName: aws.String(clusterName),
		//		})).Return(nil)
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1version, addonARN, addonStatusActive),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
		//{
		//	name: "1 installed and 0 desired - addon has status of deleting",
		//	expect: func(m *mock_eksiface.MockEKSAPIMockRecorder) {
		//		m.WaitUntilAddonDeleted(gomock.Eq(&eks.DescribeAddonInput{
		//			AddonName:   aws.String(addon1Name),
		//			ClusterName: aws.String(clusterName),
		//		})).Return(nil)
		//	},
		//	installedAddons: []*EKSAddon{
		//		createInstalledAddon(addon1Name, addon1version, addonARN, addonStatusDeleting),
		//	},
		//	expectCreateError: false,
		//	expectDoError:     false,
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			mockControl := gomock.NewController(t)
			defer mockControl.Finish()

			eksMock := mock_eksiface.NewMockEKSAPI(mockControl)
			tc.expect(eksMock.EXPECT())

			ctx := context.TODO()

			planner := NewPlan(clusterName, tc.currentIdentityProvider, tc.desiredIdentityProvider, eksMock, log)
			procedures, err := planner.Create(ctx)
			if tc.expectCreateError {
				g.Expect(err).To(HaveOccurred())
				return
			}
			g.Expect(err).To(BeNil())
			//g.Expect(procedures).NotTo(BeNil())

			for _, proc := range procedures {
				procErr := proc.Do(ctx)
				if tc.expectDoError {
					g.Expect(procErr).To(HaveOccurred())
					return
				}
				g.Expect(procErr).To(BeNil())
			}
		})
	}
}
