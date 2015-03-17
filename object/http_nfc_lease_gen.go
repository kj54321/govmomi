/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package object

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type HttpNfcLease struct {
	Common
}

func NewHttpNfcLease(c *vim25.Client, ref types.ManagedObjectReference) *HttpNfcLease {
	return &HttpNfcLease{
		Common: NewCommon(c, ref),
	}
}

// HttpNfcLeaseAbort wraps methods.HttpNfcLeaseAbort
func (o HttpNfcLease) HttpNfcLeaseAbort(ctx context.Context, fault *types.LocalizedMethodFault) error {
	req := types.HttpNfcLeaseAbort{
		This:  o.Reference(),
		Fault: fault,
	}

	_, err := methods.HttpNfcLeaseAbort(ctx, o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// HttpNfcLeaseComplete wraps methods.HttpNfcLeaseComplete
func (o HttpNfcLease) HttpNfcLeaseComplete(ctx context.Context) error {
	req := types.HttpNfcLeaseComplete{
		This: o.Reference(),
	}

	_, err := methods.HttpNfcLeaseComplete(ctx, o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// HttpNfcLeaseGetManifest wraps methods.HttpNfcLeaseGetManifest
func (o HttpNfcLease) HttpNfcLeaseGetManifest(ctx context.Context) error {
	req := types.HttpNfcLeaseGetManifest{
		This: o.Reference(),
	}

	_, err := methods.HttpNfcLeaseGetManifest(ctx, o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// HttpNfcLeaseProgress wraps methods.HttpNfcLeaseProgress
func (o HttpNfcLease) HttpNfcLeaseProgress(ctx context.Context, percent int) error {
	req := types.HttpNfcLeaseProgress{
		This:    o.Reference(),
		Percent: percent,
	}

	_, err := methods.HttpNfcLeaseProgress(ctx, o.c, &req)
	if err != nil {
		return err
	}

	return nil
}
