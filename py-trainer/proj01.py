#!/usr/bin/env python3
'''
Source: https://github.com/mkfzdmr/Deep-Learning-based-Emotion-Recognition/blob/master/train_network.py
'''
import math


# Transform Cartesian coordinates to spherical
def cart2sph(x, y, z):
    x2_y2 = x ** 2 + y ** 2
    r = math.sqrt(x2_y2 + z ** 2)  # r
    elev = math.atan2(z, math.sqrt(x2_y2))  # Elevation
    az = math.atan2(y, x)  # Azimuth
    return r, elev, az


# Transform polar coordinates to Cartesian
def pol2cart(theta, rho):
    return rho * math.cos(theta), rho * math.sin(theta)


def azim_proj(pos):
    """
    Computes the Azimuthal Equidistant Projection of input point in 3D Cartesian Coordinates.
    Imagine a plane being placed against (tangent to) a globe. If
    a light source inside the globe projects the graticule onto
    the plane the result would be a planar, or azimuthal, map
    projection.

    :param pos: position in 3D Cartesian coordinates
    :return: projected coordinates using Azimuthal Equidistant Projection
    """
    [r, elev, az] = cart2sph(pos[0], pos[1], pos[2])
    return pol2cart(az, math.pi / 2 - elev)


def test_sph():
    x, y, z = 1, 0, 1
    r, elev, az = cart2sph(x, y, z)
    # r = 1.414, elev = 0.785(rad), az = 0(rad)
    print(f'Cartesian: ({x}, {y}, {z}) -> Spherical: ({r}, {elev}, {az})')


def test_proj():
    # The reasoning:
    # As can be seen from test001(), point (1, 0, 1) has an elevation angle of 0.785 rad (45 degrees).
    # If elevation is pi/2, then the point is at the top of the sphere/our heads, thus being projected to (0, 0).
    print(azim_proj((1, 0, 1)))      # (0.785, 0)
    print(azim_proj((0, 1, 0)))      # (0, 1.571) -- indicates a circle of radius pi/2, not of the original sphere
    print(azim_proj((0, 0, 1.414)))  # (0, 0) -- correct
