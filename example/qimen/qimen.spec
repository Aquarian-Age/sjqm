#
# spec file for package groff-utf8-1.22.tar.gz
#
# Copyright (c) 2018 SUSE LINUX GmbH, Nuernberg, Germany.
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via http://bugs.opensuse.org/
#


Name:          qimen
Version:       0.6.9 
Release:       0
Summary:       QiMen
License:       MIT
Group:         Application
Url:           https://github.com/Aquarian-Age/sjqm/releases
Source0:       qimen.tar.gz 

BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
QiMen


%prep

%build
tar xzvf ../SOURCES/qimen.tar.gz -C ${RPM_BUILD_ROOT}/
cd ${RPM_BUILD_ROOT}/
go build -o qimen -ldflags="-s -w" -tags timetzdata

%install
mkdir -p ${RPM_BUILD_ROOT}/usr/local/bin/
cp -f ${RPM_BUILD_ROOT}/qimen ${RPM_BUILD_ROOT}/usr/local/bin/qimen
mkdir -p ${RPM_BUILD_ROOT}/usr/local/share/applications/
mkdir -p ${RPM_BUILD_ROOT}/usr/local/share/icons/
cp -f ${RPM_BUILD_ROOT}/qimen.svg ${RPM_BUILD_ROOT}/usr/local/share/icons/qimen.svg
cp -f ${RPM_BUILD_ROOT}/qimen.desktop ${RPM_BUILD_ROOT}/usr/local/share/applications/qimen.desktop

rm -rf ${RPM_BUILD_ROOT}/*.go
rm -rf ${RPM_BUILD_ROOT}/go.*
rm -rf ${RPM_BUILD_ROOT}/qimen*
rm -rf ${RPM_BUILD_ROOT}/qimen.svg

%files
%defattr(-,root,root)
/usr/local/bin/qimen
/usr/local/share/applications/qimen.desktop
/usr/local/share/icons/qimen.svg


%changelog
