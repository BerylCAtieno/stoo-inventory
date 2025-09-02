import { useState } from 'react'

import { 
  User, 
  Package, 
  ShoppingCart, 
  BarChart3, 
  Users, 
  FileText, 
  Settings,
  Menu,
  X,
  Archive
} from 'lucide-react';

function SideBar() {
    
const Sidebar = ({ activeTab, setActiveTab, userRole }) => {
  const [isCollapsed, setIsCollapsed] = useState(false);

  const menuItems = [
    { id: 'dashboard', label: 'Dashboard', icon: BarChart3, roles: ['Admin', 'Manager', 'Staff'] },
    { id: 'products', label: 'Products', icon: Package, roles: ['Admin', 'Manager'] },
    { id: 'categories', label: 'Categories', icon: Archive, roles: ['Admin', 'Manager'] },
    { id: 'orders', label: 'Orders', icon: ShoppingCart, roles: ['Admin', 'Manager', 'Staff'] },
    { id: 'users', label: 'Users', icon: Users, roles: ['Admin'] },
    { id: 'reports', label: 'Reports', icon: FileText, roles: ['Admin', 'Manager'] },
    { id: 'audit', label: 'Audit Logs', icon: Settings, roles: ['Admin', 'Manager'] }
  ];

  const filteredItems = menuItems.filter(item => item.roles.includes(userRole));

  return (
    <aside className={`bg-purple-900 text-white transition-all duration-300 ${isCollapsed ? 'w-16' : 'w-64'}`}>
      <div className="p-6">
        <div className="flex items-center justify-between">
          {!isCollapsed && (
            <div className="flex items-center space-x-3">
              <Package className="w-8 h-8 text-purple-300" />
              <h2 className="text-xl font-bold">Inventory</h2>
            </div>
          )}
          <button
            onClick={() => setIsCollapsed(!isCollapsed)}
            className="p-2 rounded-lg hover:bg-purple-800"
          >
            {isCollapsed ? <Menu className="w-5 h-5" /> : <X className="w-5 h-5" />}
          </button>
        </div>
      </div>

      <nav className="px-4 pb-6">
        <ul className="space-y-2">
          {filteredItems.map((item) => (
            <li key={item.id}>
              <button
                onClick={() => setActiveTab(item.id)}
                className={`w-full flex items-center space-x-3 px-3 py-2 rounded-lg transition-colors duration-200 ${
                  activeTab === item.id
                    ? 'bg-purple-700 text-white'
                    : 'text-purple-200 hover:bg-purple-800 hover:text-white'
                }`}
              >
                <item.icon className="w-5 h-5" />
                {!isCollapsed && <span>{item.label}</span>}
              </button>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  );
};

}

export default SideBar