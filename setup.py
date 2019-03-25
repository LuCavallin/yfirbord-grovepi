# -*- coding: utf-8 -*-

from setuptools import setup, find_packages


with open('README.rst') as f:
    readme = f.read()

with open('LICENSE.txt') as f:
    license = f.read()

setup(
    name='hytta',
    version='1.0.0',
    description='Gather and view sensors data from the RaspberryPi.',
    long_description=readme,
    author='Luca Cavallin',
    author_email='me@lucavall.in',
    url='https://lcv.sh',
    license=license,
    packages=find_packages(exclude=('tests'))
)
