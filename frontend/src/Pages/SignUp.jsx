import React, { useState } from 'react';
import { Eye, EyeOff, User, Mail, Phone, Lock } from 'lucide-react';

// Placeholder for your existing sign-up button component
const SignUpButton = ({ onClick, disabled, children }) => (
  <button
    onClick={onClick}
    disabled={disabled}
    className="w-full bg-indigo-900 hover:bg-indigo-800 disabled:bg-gray-400 text-white font-medium py-3 px-4 rounded-lg transition-colors duration-200"
  >
    {children}
  </button>
);

const RegistrationPage = () => {
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    email: '',
    phoneNo: '',
    password: '',
    agreeToTerms: false
  });
  
  const [showPassword, setShowPassword] = useState(false);
  const [errors, setErrors] = useState({});

  const handleInputChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
    
    // Clear error when user starts typing
    if (errors[name]) {
      setErrors(prev => ({
        ...prev,
        [name]: ''
      }));
    }
  };

  const validateForm = () => {
    const newErrors = {};
    
    if (!formData.firstName.trim()) newErrors.firstName = 'First name is required';
    if (!formData.lastName.trim()) newErrors.lastName = 'Last name is required';
    if (!formData.email.trim()) newErrors.email = 'Email is required';
    else if (!/\S+@\S+\.\S+/.test(formData.email)) newErrors.email = 'Email is invalid';
    if (!formData.phoneNo.trim()) newErrors.phoneNo = 'Phone number is required';
    if (!formData.password) newErrors.password = 'Password is required';
    else if (formData.password.length < 8) newErrors.password = 'Password must be at least 8 characters';
    if (!formData.agreeToTerms) newErrors.agreeToTerms = 'You must agree to terms and conditions';
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSignUp = () => {
    if (validateForm()) {
      console.log('Registration data:', formData);
      // Handle registration logic here
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex">
      {/* Left side - Illustration */}
      <div className="hidden lg:flex lg:w-1/2 bg-gradient-to-br from-blue-50 to-indigo-100 items-center justify-center p-12">
        <div className="max-w-lg">
          {/* Isometric Illustration */}
          <div className="relative mb-8">
            <svg
              viewBox="0 0 400 300"
              className="w-full h-auto max-w-md mx-auto"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              {/* Base platform */}
              <path
                d="M80 220 L320 220 L360 180 L120 180 Z"
                fill="#E5E7EB"
                stroke="#D1D5DB"
                strokeWidth="1"
              />
              
              {/* Left building */}
              <g>
                <path d="M120 180 L120 120 L160 100 L160 160 Z" fill="#3B82F6" />
                <path d="M160 160 L160 100 L200 120 L200 180 Z" fill="#1E40AF" />
                <path d="M120 180 L160 160 L200 180 L160 200 Z" fill="#1D4ED8" />
              </g>
              
              {/* Center charts/graphs */}
              <g>
                <path d="M220 180 L220 140 L260 120 L260 160 Z" fill="#F59E0B" />
                <path d="M260 160 L260 120 L300 140 L300 180 Z" fill="#D97706" />
                <path d="M220 180 L260 160 L300 180 L260 200 Z" fill="#F59E0B" />
                
                {/* Chart elements */}
                <rect x="230" y="150" width="8" height="20" fill="#FFF" />
                <rect x="242" y="145" width="8" height="25" fill="#FFF" />
                <rect x="254" y="140" width="8" height="30" fill="#FFF" />
              </g>
              
              {/* Right building */}
              <g>
                <path d="M320 180 L320 130 L360 110 L360 150 Z" fill="#EF4444" />
                <path d="M360 150 L360 110 L400 130 L400 170 Z" fill="#DC2626" />
                <path d="M320 180 L360 150 L400 170 L360 190 Z" fill="#EF4444" />
              </g>
              
              {/* Floating elements */}
              <circle cx="100" cy="100" r="12" fill="#10B981" opacity="0.8" />
              <circle cx="380" cy="90" r="8" fill="#F59E0B" opacity="0.8" />
              <circle cx="200" cy="80" r="6" fill="#8B5CF6" opacity="0.8" />
              
              {/* Person figure */}
              <g transform="translate(180, 180)">
                <circle cx="0" cy="-25" r="8" fill="#FCA5A5" />
                <rect x="-6" y="-20" width="12" height="20" rx="2" fill="#3B82F6" />
                <rect x="-4" y="0" width="8" height="15" rx="1" fill="#1F2937" />
              </g>
            </svg>
          </div>
        </div>
      </div>

      {/* Right side - Registration Form */}
      <div className="flex-1 flex items-center justify-center p-8">
        <div className="w-full max-w-md">
          {/* Header */}
          <div className="text-center mb-8">
            <div className="w-12 h-12 bg-cyan-400 rounded-lg mx-auto mb-4 flex items-center justify-center">
              <div className="w-6 h-6 bg-white rounded-sm"></div>
            </div>
            <h1 className="text-2xl font-bold text-gray-900 mb-2">Register</h1>
            <p className="text-gray-600 text-sm">
              Manage all your inventory efficiently
            </p>
            <p className="text-gray-600 text-sm">
              Let's get you all set up so you can verify your personal account and begin setting up your work profile.
            </p>
          </div>

          {/* Registration Form */}
          <div className="space-y-4">
            {/* Name Fields */}
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  First name
                </label>
                <div className="relative">
                  <input
                    type="text"
                    name="firstName"
                    value={formData.firstName}
                    onChange={handleInputChange}
                    placeholder="Enter your name"
                    className={`w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent ${
                      errors.firstName ? 'border-red-500' : 'border-gray-300'
                    }`}
                  />
                  <User className="absolute right-3 top-2.5 h-5 w-5 text-gray-400" />
                </div>
                {errors.firstName && (
                  <p className="text-red-500 text-xs mt-1">{errors.firstName}</p>
                )}
              </div>
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Last name
                </label>
                <input
                  type="text"
                  name="lastName"
                  value={formData.lastName}
                  onChange={handleInputChange}
                  placeholder="minimum 8 characters"
                  className={`w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent ${
                    errors.lastName ? 'border-red-500' : 'border-gray-300'
                  }`}
                />
                {errors.lastName && (
                  <p className="text-red-500 text-xs mt-1">{errors.lastName}</p>
                )}
              </div>
            </div>

            {/* Email and Phone */}
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Email
                </label>
                <div className="relative">
                  <input
                    type="email"
                    name="email"
                    value={formData.email}
                    onChange={handleInputChange}
                    placeholder="Enter your email"
                    className={`w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent ${
                      errors.email ? 'border-red-500' : 'border-gray-300'
                    }`}
                  />
                  <Mail className="absolute right-3 top-2.5 h-5 w-5 text-gray-400" />
                </div>
                {errors.email && (
                  <p className="text-red-500 text-xs mt-1">{errors.email}</p>
                )}
              </div>
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  Phone no
                </label>
                <div className="relative">
                  <input
                    type="tel"
                    name="phoneNo"
                    value={formData.phoneNo}
                    onChange={handleInputChange}
                    placeholder="minimum 8 characters"
                    className={`w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent ${
                      errors.phoneNo ? 'border-red-500' : 'border-gray-300'
                    }`}
                  />
                  <Phone className="absolute right-3 top-2.5 h-5 w-5 text-gray-400" />
                </div>
                {errors.phoneNo && (
                  <p className="text-red-500 text-xs mt-1">{errors.phoneNo}</p>
                )}
              </div>
            </div>

            {/* Password */}
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Password
              </label>
              <div className="relative">
                <input
                  type={showPassword ? 'text' : 'password'}
                  name="password"
                  value={formData.password}
                  onChange={handleInputChange}
                  placeholder="Enter your password"
                  className={`w-full px-3 py-2 pr-10 border rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent ${
                    errors.password ? 'border-red-500' : 'border-gray-300'
                  }`}
                />
                <Lock className="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute right-3 top-2.5 text-gray-400 hover:text-gray-600"
                >
                  {showPassword ? <EyeOff className="h-5 w-5" /> : <Eye className="h-5 w-5" />}
                </button>
              </div>
              {errors.password && (
                <p className="text-red-500 text-xs mt-1">{errors.password}</p>
              )}
            </div>

            {/* Terms Checkbox */}
            <div className="flex items-start">
              <input
                type="checkbox"
                name="agreeToTerms"
                checked={formData.agreeToTerms}
                onChange={handleInputChange}
                className="mt-1 h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
              />
              <label className="ml-2 text-sm text-gray-600">
                I agree to all terms, privacy policies, and fees
              </label>
            </div>
            {errors.agreeToTerms && (
              <p className="text-red-500 text-xs">{errors.agreeToTerms}</p>
            )}

            {/* Sign Up Button */}
            <div className="mt-6">
              <SignUpButton
                onClick={handleSignUp}
                disabled={!formData.agreeToTerms}
              >
                Sign up
              </SignUpButton>
            </div>

            {/* Login Link */}
            <div className="text-center mt-4">
              <span className="text-sm text-gray-600">
                Already have an account?{' '}
                <button
                  type="button"
                  className="text-indigo-600 hover:text-indigo-500 font-medium"
                  onClick={() => console.log('Navigate to login')}
                >
                  Log in
                </button>
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegistrationPage;